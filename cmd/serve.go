package cmd

import (
	"github.com/carbans/simpledns/pkg/simpledns"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve command is used to start the DNS server",
	Run: func(cmd *cobra.Command, args []string) {
		simpledns.Serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
