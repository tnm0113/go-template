package cmd

import (
	"fmt"
	"os"

	"github.com/c4i/go-template/internal/config"
	"github.com/c4i/go-template/internal/db"
	"github.com/c4i/go-template/internal/hapi"
	"github.com/c4i/go-template/internal/service"
	"github.com/spf13/cobra"
)

const probeFlag string = "probe"

var serverCmd = &cobra.Command{
	Use:   "server",
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
		fmt.Println("read config error")
		os.Exit(1)
	}
	fmt.Printf("%v", cfg)

	mongo, err := db.ConnectToMongoDB(cfg)
	if err != nil {
		fmt.Printf("Failed to connect to Mongo")
	}
	svc := service.New(mongo)

	fmt.Println("start http server")
	http_server := hapi.NewServer(svc, cfg)
	http_server.Start()
}

func runReadiness(verbose bool) {

}
