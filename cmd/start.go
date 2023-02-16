package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/c4i/go-template/internal/config"
	"github.com/c4i/go-template/internal/db"
	"github.com/c4i/go-template/internal/gapi"
	"github.com/c4i/go-template/internal/hapi"
	"github.com/c4i/go-template/internal/hapi/router"
	"github.com/c4i/go-template/internal/service"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
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

func runServer() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("read config error")
		os.Exit(1)
	}

	mongo, err := db.ConnectToMongoDB(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to MongoDB")
		os.Exit(1)
	}
	svc := service.New(mongo)

	errs := make(chan error, 2)

	log.Info().Msg("start http server")
	http_server := hapi.NewServer(svc, cfg)
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
