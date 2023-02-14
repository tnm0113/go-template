package cmd

import (
	"fmt"
	"os"

	"github.com/c4i/go-template/internal/config"
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
	// mongo := connectToMongoDB()
	fmt.Println("start server")
}

// func connectToMongoDB(config config.MongoDB) *mongo.Database {
// 	addr := fmt.Sprintf("mongodb://%s:%d/?replicaset=%s", config.Host, config.Port, config.Replica)
// 	credential := options.Credential{
// 		Username: config.Username,
// 		Password: config.Password,
// 	}
// 	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(addr).SetAuth(credential))
// 	if err != nil {
// 		os.Exit(1)
// 	}

// 	return client.Database(config.DbName)
// }

func runReadiness(verbose bool) {

}
