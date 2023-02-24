package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"192.168.205.151/vq2-go/go-template/internal/config"
	"192.168.205.151/vq2-go/go-template/internal/gapi"
	"192.168.205.151/vq2-go/go-template/internal/hapi"
	"192.168.205.151/vq2-go/go-template/internal/hapi/router"
	"192.168.205.151/vq2-go/go-template/internal/mq"
	"192.168.205.151/vq2-go/go-template/internal/service"
	"github.com/qiniu/qmgo"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

const probeFlag string = "probe"

var serverCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the server",
	Long:  "Starts server",
	Run: func(cmd *cobra.Command, args []string) {
		probeReadiness, err := cmd.Flags().GetBool(probeFlag)
		if err != nil {
			fmt.Printf("Error while parsing flags: %v\n", err)
			os.Exit(1)
		}

		if probeReadiness {
			runReadiness(true)
		}

		runServer()
	},
}

func init() {
	serverCmd.Flags().BoolP(probeFlag, "p", false, "Probe readiness before startup.")
	rootCmd.AddCommand(serverCmd)
}

func initTracer(cfg config.OtherConfig) (*sdktrace.TracerProvider, error) {
	collectorUrl := fmt.Sprintf("http://%s:%d/api/v2/spans", cfg.TracingHost, cfg.TracingPort)
	exporter, err := zipkin.New(collectorUrl)
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(config.ModuleName),
		)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}

func runServer() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("read config error")
		os.Exit(1)
	}

	log.Debug().Msgf("config %v", cfg)

	if cfg.OtherConfig.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05 02-01-2006"})
	}

	tp, err := initTracer(cfg.OtherConfig)
	if err != nil {
		log.Error().Err(err).Msg("Init tracer error")
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Error().Msgf("Error shutting down tracer provider: %v", err)
		}
	}()

	ctx := context.Background()
	addr := fmt.Sprintf("mongodb://%s:%d/?replicaset=%s", cfg.DbConfig.DBHost, cfg.DbConfig.DBPort, cfg.DbConfig.DBReplica)
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: addr})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to MongoDB")
		os.Exit(1)
	}
	db := client.Database(cfg.DbConfig.DBName)

	rabbit := mq.New(cfg.RabbitmqConfig)
	err = rabbit.Connect()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to rabbit")
		os.Exit(1)
	}
	defer rabbit.Shutdown()

	err = rabbit.Declare()
	if err != nil {
		log.Error().Err(err).Msg("Failed to setup exchange and queue")
	}
	err = rabbit.Subcribe("user.*")
	if err != nil {
		log.Error().Err(err).Msg("Failed to subcribe")
	}

	svc := service.New(db, rabbit)

	errs := make(chan error, 2)

	log.Info().Msg("start http server")
	http_server := hapi.NewServer(svc, cfg)
	http_server.InitI18n()
	router.Init(http_server)
	go http_server.Start(errs)

	log.Info().Msg("start grpc server")
	grpc_server := gapi.NewServer(cfg, svc)
	go grpc_server.Start(errs)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	err = <-errs
	log.Fatal().Err(err).Msg("Services terminate")
}

func runReadiness(verbose bool) {

}
