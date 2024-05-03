package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "simpledns",
	Short: "SimpleDNS is a simple DNS server to resolve domain names to IP addresses",
	Long: `SimpleDNS It is a simple implementation of a DNS server that can be used to resolve domain names to IP addresses.
	Not production ready, but can be used for educational purposes.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.simpledns.yaml)")
}
