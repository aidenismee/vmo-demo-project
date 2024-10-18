package cmd

import (
	"fmt"
	"github.com/nekizz/vmo-demo-project/pkg/file"
	"log"

	"github.com/nekizz/vmo-demo-project/internal/contract"
	fileProcessor "github.com/nekizz/vmo-demo-project/internal/file_processor"
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
	filerSvc := file.NewService()
	fileProcesserSvc := fileProcessor.NewService(nil, filerSvc)
	fileProcesserHandler := fileProcessor.NewHandler(fileProcesserSvc)

	count, err := fileProcesserHandler.CountFileLine(&contract.CountLineInput{Path: filePath})
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(count)
}
