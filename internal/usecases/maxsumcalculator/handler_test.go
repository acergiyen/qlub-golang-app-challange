package maxsumcalculator_test

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/acergiyen/qlub-golang-app-challange/internal/usecases/maxsumcalculator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHandlers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MaxSumCalculator Suite")
}

var _ = Describe("MaxSumCalculator", func() {
	var (
		calculator *maxsumcalculator.MaxSumCalculator
		treeData   maxsumcalculator.TreeRequest
		maxSumPath int
	)
	var buff bytes.Buffer
	lgr := log.New(&buff, "", log.LstdFlags)
	calculator = maxsumcalculator.NewMaxSumCalculator(lgr)

	Context("MaxSumCalculator", func() {
		When("data comes from testcase1", func() {
			BeforeEach(func() {
				jsonData, _ := os.ReadFile("testdata/testcase1.json")
				json.Unmarshal(jsonData, &treeData)
				maxSumPath = calculator.Handle(treeData)
			})

			It("should be 18", func() {
				Expect(18).To(Equal(maxSumPath))
			})
		})
		When("data comes from testcase2", func() {
			BeforeEach(func() {
				jsonData, _ := os.ReadFile("testdata/testcase2.json")
				json.Unmarshal(jsonData, &treeData)
				maxSumPath = calculator.Handle(treeData)
			})

			It("should be 0", func() {
				Expect(6).To(Equal(maxSumPath))
			})
		})
		When("data comes from testcase3", func() {
			BeforeEach(func() {
				jsonData, _ := os.ReadFile("testdata/testcase3.json")
				json.Unmarshal(jsonData, &treeData)
				maxSumPath = calculator.Handle(treeData)
			})

			It("should be 0", func() {
				Expect(159).To(Equal(maxSumPath))
			})
		})
	})
})
