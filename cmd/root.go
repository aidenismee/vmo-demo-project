package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "futil",
		Short: "futil service root command",
		Long:  `File Utility`,
	}
)

func init() {
	cobra.OnInitialize(initConfig)
}

// TODO: init config from .env
func initConfig() {}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
