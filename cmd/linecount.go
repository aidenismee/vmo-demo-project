package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/nekizz/vmo-demo-project/internal/contract"
	fileProcessor "github.com/nekizz/vmo-demo-project/internal/file_processor"
	filerPkg "github.com/nekizz/vmo-demo-project/pkg/filer"
)

func init() {
	linecountCmd.Flags().StringVarP(&filePath, "file", "f", "", "the input file")

	rootCmd.AddCommand(linecountCmd)
}

var linecountCmd = &cobra.Command{
	Use:   "linecount",
	Short: "Print line count of file",
	Run:   lineCounterHandler,
}

func lineCounterHandler(cmd *cobra.Command, args []string) {
	filer := filerPkg.NewFiler()
	fileProcessorSvc := fileProcessor.NewService(nil, filer)
	fileProcessorHandler := fileProcessor.NewHandler(fileProcessorSvc)

	result, err := fileProcessorHandler.CountFileLine(&contract.CountLineInput{Path: filePath})
	if err != nil {
		log.Fatalf("Err: %s", err)
	}

	log.Println(result)
}
