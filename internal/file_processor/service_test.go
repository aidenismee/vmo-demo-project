package file_processor

import (
	"errors"
	filePkg "github.com/nekizz/vmo-demo-project/pkg/file"
	"github.com/nekizz/vmo-demo-project/pkg/hasher"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
	"os"
	"testing"
)

func TestFileProcessorService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "File Processor Service Suite")
}

var _ = Describe("File Processor Service Feature", func() {
	var (
		ctrl          *gomock.Controller
		target        *service
		filerService  *filePkg.MockService
		hasherService *hasher.MockService
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		filerService = filePkg.NewMockService(ctrl)
		hasherService = hasher.NewMockService(ctrl)
		target = NewService(hasherService, filerService)
	})

	Describe("Count Line Test", func() {
		When("count line successfully", func() {
			It("count line success, should return nil", func() {
				input := "myfile.txt"
				expectLine := 5

				filerService.EXPECT().OpenFile(gomock.Any()).
					DoAndReturn(func(arg0 string) (*os.File, error) {
						return new(os.File), nil
					})

				filerService.EXPECT().CountLine(gomock.Any()).
					DoAndReturn(func(arg0 *os.File) (int, error) {
						return 5, nil
					})

				resp, err := target.CountFileLine(input)
				Expect(err).To(BeNil())
				Expect(resp).To(Equal(expectLine))
			})
		})

		When("count line failed", func() {
			It("open file failed, should return error", func() {
				input := "myfile.txt"
				expectLine := 0

				filerService.EXPECT().OpenFile(gomock.Any()).
					DoAndReturn(func(arg0 string) (*os.File, error) {
						return nil, errors.New("file not exist")
					})

				resp, err := target.CountFileLine(input)
				Expect(err).NotTo(BeNil())
				Expect(resp).To(Equal(expectLine))
			})

			It("count failed, should return error", func() {
				input := "myfile.txt"
				expectLine := 0

				filerService.EXPECT().OpenFile(gomock.Any()).
					DoAndReturn(func(arg0 string) (*os.File, error) {
						return nil, nil
					})

				filerService.EXPECT().CountLine(gomock.Any()).
					DoAndReturn(func(arg0 *os.File) (int, error) {
						return 0, errors.New("can't count file")
					})

				resp, err := target.CountFileLine(input)
				Expect(err).NotTo(BeNil())
				Expect(resp).To(Equal(expectLine))
			})
		})
	})

	Describe("Check Sum Test", func() {
		When("check sum successfully", func() {
			It("check sum success, should return nil", func() {
				input := "myfile.txt"
				expectLine := "abc"

				filerService.EXPECT().OpenFile(gomock.Any()).
					DoAndReturn(func(arg0 string) (*os.File, error) {
						return new(os.File), nil
					})

				hasherService.EXPECT().HashFile(gomock.Any()).
					DoAndReturn(func(arg0 *os.File) (string, error) {
						return "abc", nil
					})

				resp, err := target.CheckSum(input)
				Expect(err).To(BeNil())
				Expect(resp).To(Equal(expectLine))
			})
		})

		When("check sum failed", func() {
			It("check sum failed, should return error", func() {
				input := "myfile.txt"
				expectLine := ""

				filerService.EXPECT().OpenFile(gomock.Any()).
					DoAndReturn(func(arg0 string) (*os.File, error) {
						return nil, errors.New("file not exist")
					})

				resp, err := target.CheckSum(input)
				Expect(err).NotTo(BeNil())
				Expect(resp).To(Equal(expectLine))
			})

			It("count failed, should return error", func() {
				input := "myfile.txt"
				expectLine := ""

				filerService.EXPECT().OpenFile(gomock.Any()).
					DoAndReturn(func(arg0 string) (*os.File, error) {
						return nil, nil
					})

				hasherService.EXPECT().HashFile(gomock.Any()).
					DoAndReturn(func(arg0 *os.File) (string, error) {
						return "", errors.New("can't count file")
					})

				resp, err := target.CheckSum(input)
				Expect(err).NotTo(BeNil())
				Expect(resp).To(Equal(expectLine))
			})
		})
	})

	AfterEach(func() {
		ctrl.Finish()
	})
})
