package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(linecountCmd)
}

var linecountCmd = &cobra.Command{
	Use:   "linecount",
	Short: "Print line count of file",
	Run:   func(cmd *cobra.Command, args []string) {},
}
