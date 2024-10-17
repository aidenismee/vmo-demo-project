package cmd

import (
	"fmt"
	"github.com/nekizz/vmo-demo-project/internal/contract"
	fileProcesser "github.com/nekizz/vmo-demo-project/internal/file-processer"
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
	fileProcesserSvc := fileProcesser.NewService()
	fileProcesserHandler := fileProcesser.NewHandler(fileProcesserSvc)

	sum, err := fileProcesserHandler.CheckSum(&contract.CheckSumInput{Path: filePath, Algo: algo})
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

	return nil
}
