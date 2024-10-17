package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
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
		log.Println(err)
		os.Exit(1)
	}
}
