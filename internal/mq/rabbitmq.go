package mq

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/c4i/go-template/internal/config"
	"github.com/pkg/errors"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	mux                  sync.RWMutex
	config               config.RabbitMQConfig
	connection           *amqp091.Connection
	dialConfig           amqp091.Config
	ChannelNotifyTimeout time.Duration
}

func New(config config.RabbitMQConfig) *RabbitMQ {
	return &RabbitMQ{
		config:               config,
		dialConfig:           amqp091.Config{Properties: amqp091.Table{"connection_name": config.ConnectionName}},
		ChannelNotifyTimeout: time.Duration(config.ChannelNotifyTimeout),
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
		log.Println("CRITICAL: Connection dropped, reconnecting")

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
				log.Println("INFO: Reconnected")

				goto WATCH
			}

			time.Sleep(time.Duration(r.config.ReconnectInterval))
		}

		log.Println(errors.Wrap(err, "CRITICAL: Failed to reconnect"))
	} else {
		log.Println("INFO: Connection dropped normally, will not reconnect")
	}
}
