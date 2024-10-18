package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	filePath string

	rootCmd = &cobra.Command{
		Use:   "futil",
		Short: "futil service root command",
		Long:  `File Utility`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
