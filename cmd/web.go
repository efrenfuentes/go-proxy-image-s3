package cmd

import (
	"go-proxy-image-s3/internal/app"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var httpAddr string

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Starts the go-proxy-image-s3 server (default to 127.0.0.1:8080)",
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize a new logger which writes messages to the standard out stream,
		// prefixed with the current date and time.

		log := log.New()

		config := app.Config{
			HttpAddr: httpAddr,
		}

		app := app.NewApp(config, log)

		err := app.Run()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(webCmd)

	webCmd.PersistentFlags().StringVar(&httpAddr, "addr", "127.0.0.1:8080", "HTTP server address")
}
