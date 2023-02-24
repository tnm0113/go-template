package mq

import (
	"context"
	"fmt"
	"sync"
	"time"

	"192.168.205.151/vq2-go/go-template/internal/config"
	"github.com/pkg/errors"
	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

const EXCHANGE_NAME string = "system-event"
const EXCHANGE_TYPE string = "topic"

type RabbitMQ struct {
	mux                  sync.RWMutex
	config               config.RabbitMQConfig
	connection           *amqp091.Connection
	dialConfig           amqp091.Config
	ChannelNotifyTimeout time.Duration
}

func New(cfg config.RabbitMQConfig) *RabbitMQ {
	moduleName := config.GetModuleName()
	return &RabbitMQ{
		config:               cfg,
		dialConfig:           amqp091.Config{Properties: amqp091.Table{"connection_name": moduleName}},
		ChannelNotifyTimeout: time.Duration(cfg.ChannelNotifyTimeout),
	}
}

// Connect creates a new connection. Use once at application
// startup.
func (r *RabbitMQ) Connect() error {
	con, err := amqp091.DialConfig(fmt.Sprintf(
		"%s://%s:%s@%s:%d/%s",
		r.config.Schema,
		r.config.Username,
		r.config.Password,
		r.config.Host,
		r.config.Port,
		r.config.Vhost,
	), r.dialConfig)
	if err != nil {
		return err
	}

	r.connection = con

	go r.reconnect()

	return nil
}

// Channel returns a new `*amqp.Channel` instance. You must
// call `defer channel.Close()` as soon as you obtain one.
// Sometimes the connection might be closed unintentionally so
// as a graceful handling, try to connect only once.
func (r *RabbitMQ) Channel() (*amqp091.Channel, error) {
	if r.connection == nil {
		if err := r.Connect(); err != nil {
			return nil, errors.New("connection is not open")
		}
	}

	channel, err := r.connection.Channel()
	if err != nil {
		return nil, err
	}

	return channel, nil
}

// Connection exposes the essentials of the current connection.
// You should not normally use this but it is there for special
// use cases.
func (r *RabbitMQ) Connection() *amqp091.Connection {
	return r.connection
}

// Shutdown triggers a normal shutdown. Use this when you wish
// to shutdown your current connection or if you are shutting
// down the application.
func (r *RabbitMQ) Shutdown() error {
	if r.connection != nil {
		return r.connection.Close()
	}

	return nil
}

// reconnect reconnects to server if the connection or a channel
// is closed unexpectedly. Normal shutdown is ignored. It tries
// maximum of 7200 times and sleeps half a second in between
// each try which equals to 1 hour.
func (r *RabbitMQ) reconnect() {
WATCH:

	conErr := <-r.connection.NotifyClose(make(chan *amqp091.Error))
	if conErr != nil {
		log.Error().Msg("CRITICAL: Connection dropped, reconnecting")

		var err error

		for i := 1; i <= r.config.ReconnectMaxAttempt; i++ {
			r.mux.RLock()
			r.connection, err = amqp091.DialConfig(fmt.Sprintf(
				"%s://%s:%s@%s:%d/%s",
				r.config.Schema,
				r.config.Username,
				r.config.Password,
				r.config.Host,
				r.config.Port,
				r.config.Vhost,
			), r.dialConfig)
			r.mux.RUnlock()

			if err == nil {
				log.Info().Msg("Reconnected")

				goto WATCH
			}

			time.Sleep(time.Duration(r.config.ReconnectInterval))
		}

		log.Error().Err(err).Msg("CRITICAL: Failed to reconnect")
	} else {
		log.Info().Msg("Connection dropped normally, will not reconnect")
	}
}

func (r *RabbitMQ) Declare() error {
	moduleName := config.GetModuleName()
	channel, err := r.Channel()
	if err != nil {
		return err
	}
	if err := channel.ExchangeDeclare(
		EXCHANGE_NAME,
		EXCHANGE_TYPE,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return errors.Wrap(err, "failed to declare exchange")
	}

	if _, err := channel.QueueDeclare(
		moduleName,
		true,
		false,
		false,
		false,
		amqp091.Table{"x-queue-mode": "lazy"},
	); err != nil {
		return errors.Wrap(err, "failed to declare queue")
	}
	return nil
}

func (r *RabbitMQ) Publish(message, topic string) error {
	channel, err := r.Channel()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get channel")
		return err
	}
	defer channel.Close()

	if err := channel.Confirm(false); err != nil {
		return errors.Wrap(err, "failed to put channel in confirmation mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := channel.PublishWithContext(
		ctx,
		EXCHANGE_NAME,
		topic,
		true,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	); err != nil {
		return errors.Wrap(err, "failed to publish message")
	}

	select {
	case ntf := <-channel.NotifyPublish(make(chan amqp091.Confirmation, 1)):
		if !ntf.Ack {
			return errors.New("failed to deliver message to exchange/queue")
		}
	case <-channel.NotifyReturn(make(chan amqp091.Return)):
		return errors.New("failed to deliver message to exchange/queue")
	case <-time.After(time.Duration(r.config.ChannelNotifyTimeout)):
		log.Debug().Msg("message delivery confirmation to exchange/queue timed out")
	}

	return nil
}

func (r *RabbitMQ) Subcribe(topic string) error {
	moduleName := config.GetModuleName()
	channel, err := r.Channel()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get channel")
		return err
	}
	if err := channel.QueueBind(
		moduleName,
		topic,
		EXCHANGE_NAME,
		false,
		nil,
	); err != nil {
		return errors.Wrap(err, "failed to bind queue")
	}
	return nil
}
