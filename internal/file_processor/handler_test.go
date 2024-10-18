package file_processor

import (
	"errors"
	"github.com/nekizz/vmo-demo-project/internal/contract"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestFileProcessorHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "File Processor Feature Suite")
}

var _ = Describe("File Processor Handler Feature", func() {
	var (
		ctrl                  *gomock.Controller
		target                *Handler
		filerProcessorService *MockService
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		filerProcessorService = NewMockService(ctrl)
		target = NewHandler(filerProcessorService)
	})

	Describe("Count Line Test", func() {
		When("count line successfully", func() {
			It("count line success, should return nil", func() {
				input := &contract.CountLineInput{Path: "myfile.txt"}
				expectLine := 5

				filerProcessorService.EXPECT().CountFileLine(gomock.Any()).
					DoAndReturn(func(arg0 string) (int, error) {
						return 5, nil
					})

				resp, err := target.CountFileLine(input)
				Expect(err).To(BeNil())
				Expect(resp).To(Equal(expectLine))
			})
		})

		When("count line failed", func() {
			It("open file failed, should return error", func() {
				input := &contract.CountLineInput{Path: ""}
				expectLine := 0

				resp, err := target.CountFileLine(input)
				Expect(err).NotTo(BeNil())
				Expect(resp).To(Equal(expectLine))
			})

			It("count failed, should return error", func() {
				input := &contract.CountLineInput{Path: "mytest.txt"}
				expectLine := 0

				filerProcessorService.EXPECT().CountFileLine(gomock.Any()).
					DoAndReturn(func(arg0 string) (int, error) {
						return 0, errors.New("some error")
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
				input := &contract.CheckSumInput{Path: "myfile.txt"}
				expectLine := "abc"

				filerProcessorService.EXPECT().CheckSum(gomock.Any()).
					DoAndReturn(func(arg0 string) (string, error) {
						return "abc", nil
					})

				resp, err := target.CheckSum(input)
				Expect(err).To(BeNil())
				Expect(resp).To(Equal(expectLine))
			})
		})

		When("check sum failed", func() {
			It("check sum failed, should return error", func() {
				input := &contract.CheckSumInput{Path: ""}
				expectLine := ""

				resp, err := target.CheckSum(input)
				Expect(err).NotTo(BeNil())
				Expect(resp).To(Equal(expectLine))
			})

			It("count failed, should return error", func() {
				input := &contract.CheckSumInput{Path: "myfile.txt"}
				expectLine := ""

				filerProcessorService.EXPECT().CheckSum(gomock.Any()).
					DoAndReturn(func(arg0 string) (string, error) {
						return "", errors.New("error occur")
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
