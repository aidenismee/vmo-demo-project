package cmd

import (
	"github.com/nekizz/vmo-demo-project/internal/contract"
	"github.com/nekizz/vmo-demo-project/internal/enum"
	fileProcessor "github.com/nekizz/vmo-demo-project/internal/file_processor"
	filerPkg "github.com/nekizz/vmo-demo-project/pkg/filer"
	hasherPkg "github.com/nekizz/vmo-demo-project/pkg/hasher"
	appError "github.com/nekizz/vmo-demo-project/pkg/utils/app_error"
	arrayUtils "github.com/nekizz/vmo-demo-project/pkg/utils/array"
	"github.com/spf13/cobra"
	"log"
)

var (
	sha1, md5, sha256 bool
	algo              string
)

func init() {
	checksumCmd.Flags().StringVarP(&filePath, "file", "f", "", "the input file")
	checksumCmd.Flags().BoolVar(&sha1, "sha1", false, "Use the SHA1 hashing algorithm")
	checksumCmd.Flags().BoolVar(&sha256, "sha256", false, "Use the SHA256 hashing algorithm")
	checksumCmd.Flags().BoolVar(&md5, "md5", false, "Use the MD5 hashing algorithm")

	rootCmd.AddCommand(checksumCmd)
}

var checksumCmd = &cobra.Command{
	Use:    "checksum",
	Short:  "Print checksum of file",
	PreRun: validateFlag,
	Run:    checkSumHandler,
}

func checkSumHandler(cmd *cobra.Command, args []string) {
	filer := filerPkg.NewFiler()
	hasher := hasherPkg.NewHasher(algo)
	fileProcessorSvc := fileProcessor.NewService(hasher, filer)
	fileProcessorHandler := fileProcessor.NewHandler(fileProcessorSvc)

	result, err := fileProcessorHandler.CheckSum(&contract.CheckSumInput{Path: filePath})
	if err != nil {
		log.Fatalf("Err: %s", err.Error())
	}

	log.Println(result)
}

func validateFlag(cmd *cobra.Command, args []string) {
	mapping := make(map[string]bool)
	mapping["sha1"] = sha1
	mapping["md5"] = md5
	mapping["sha256"] = sha256

	count := 0
	for key, value := range mapping {
		if value {
			algo = key
			count++
		}
	}

	if count == 0 {
		log.Fatalf("Err: %s", appError.ErrEmptyAlgorithm)
	}

	if count > 1 {
		log.Fatalf("Err: %s", appError.ErrMoreThanOneAlgorithm)
	}

	if !arrayUtils.Contains([]enum.Algorithm{enum.Sha256, enum.Sha1, enum.Md5},
		enum.Algorithm(algo)) {
		log.Fatalf("Err: %s", appError.ErrInvalidAlgorithm)
	}
}
