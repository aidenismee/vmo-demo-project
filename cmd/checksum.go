package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checksumCmd)
}

var checksumCmd = &cobra.Command{
	Use:   "checksum",
	Short: "Print checksum of file",
	Run:   func(cmd *cobra.Command, args []string) {},
}
