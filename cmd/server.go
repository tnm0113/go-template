package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
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
}

func runServer() {

}

func runReadiness(verbose bool) {

}
