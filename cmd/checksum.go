package cmd

import (
	"fmt"
	"github.com/nekizz/vmo-demo-project/internal/contract"
	"github.com/nekizz/vmo-demo-project/internal/enum"
	fileProcessor "github.com/nekizz/vmo-demo-project/internal/file_processor"
	"github.com/nekizz/vmo-demo-project/pkg/file"
	"github.com/nekizz/vmo-demo-project/pkg/hasher"
	arrayUtils "github.com/nekizz/vmo-demo-project/pkg/utils/array"
	"github.com/spf13/cobra"
)

var (
	sha1, md5, sha256 bool
	algo              string
)

func init() {
	checksumCmd.Flags().StringVarP(&filePath, "file", "f", "", "file location")
	checksumCmd.Flags().BoolVar(&sha1, "sha1", false, "Use the SHA1 hashing algorithm")
	checksumCmd.Flags().BoolVar(&sha256, "sha256", false, "Use the SHA256 hashing algorithm")
	checksumCmd.Flags().BoolVar(&md5, "md5", false, "Use the MD5 hashing algorithm")

	rootCmd.AddCommand(checksumCmd)
}

var checksumCmd = &cobra.Command{
	Use:   "checksum",
	Short: "Print checksum of file",
	Args:  validateFlag,
	Run:   checkSumHandler,
}

func checkSumHandler(cmd *cobra.Command, args []string) {
	hasherSvc := hasher.NewService(algo)
	filerSvc := file.NewService()
	fileProcesserSvc := fileProcessor.NewService(hasherSvc, filerSvc)
	fileProcesserHandler := fileProcessor.NewHandler(fileProcesserSvc)

	sum, err := fileProcesserHandler.CheckSum(&contract.CheckSumInput{Path: filePath})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(sum)
}

func validateFlag(cmd *cobra.Command, args []string) error {
	mapping := make(map[string]bool)
	mapping["sha1"] = sha1
	mapping["md5"] = md5
	mapping["sha256"] = sha256

	count := 0
	for key, value := range mapping {
		if count > 1 {
			return fmt.Errorf("invalid flag")
		}
		if value {
			algo = key
			count++
		}
	}

	if !arrayUtils.Contains([]enum.Algorithm{enum.Sha256, enum.Sha1, enum.Md5}, enum.Algorithm(algo)) {
		return fmt.Errorf("invalid algorithm")
	}

	return nil
}
