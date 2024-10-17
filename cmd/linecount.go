package cmd

import (
	"fmt"
	"log"

	"github.com/nekizz/vmo-demo-project/internal/contract"
	fileProcesser "github.com/nekizz/vmo-demo-project/internal/file-processer"
	"github.com/spf13/cobra"
)

func init() {
	linecountCmd.Flags().StringVarP(&filePath, "file", "f", "", "file location")

	rootCmd.AddCommand(linecountCmd)
}

var linecountCmd = &cobra.Command{
	Use:   "linecount",
	Short: "Print line count of file",
	Run:   lineCounterHandler,
}

func lineCounterHandler(cmd *cobra.Command, args []string) {
	fileProcesserSvc := fileProcesser.NewService()
	fileProcesserHandler := fileProcesser.NewHandler(fileProcesserSvc)

	count, err := fileProcesserHandler.CountFileLine(&contract.CountLineInput{Path: filePath})
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(count)
}
