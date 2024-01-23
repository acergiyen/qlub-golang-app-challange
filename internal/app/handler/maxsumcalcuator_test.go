package handler_test

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/acergiyen/qlub-golang-app-challange/internal/app/handler"
	"github.com/acergiyen/qlub-golang-app-challange/internal/usecases/maxsumcalculator"
	"github.com/gin-gonic/gin"
)

func TestHandlers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handler Suite")
}

var _ = Describe("Handler", func() {
	var (
		calculator *maxsumcalculator.MaxSumCalculator
		h          *handler.Handler
		c          *gin.Context
	)

	BeforeEach(func() {
		var buff bytes.Buffer
		lgr := log.New(&buff, "", log.LstdFlags)
		calculator = maxsumcalculator.NewMaxSumCalculator(lgr)
		h = handler.NewHandler(calculator)
		c, _ = gin.CreateTestContext(httptest.NewRecorder())
	})

	Describe("MaxSumCalculator", func() {
		Context("with valid JSON input", func() {
			It("should return maxPathSum in the response", func() {
				jsonRequest := `{
					"tree": {
					"nodes": [
					{"id": "1", "left": "2", "right": "3", "value": 1},
					{"id": "3", "left": null, "right": null, "value": 3},
					{"id": "2", "left": null, "right": null, "value": 2}
					],
					"root": "1"
					}
				}`

				c.Request, _ = http.NewRequest(http.MethodPost, "/maxsumcalculator", StringReader(jsonRequest))
				c.Request.Header.Set("Content-Type", "application/json")

				h.MaxSumCalculator(c)

				Expect(c.Writer.Status()).To(Equal(http.StatusOK))
			})
		})

		Context("with invalid JSON input", func() {
			It("should return a 400 Bad Request response", func() {
				invalidJSON := "invalid JSON"

				c.Request, _ = http.NewRequest(http.MethodPost, "/maxsumcalculator", StringReader(invalidJSON))
				c.Request.Header.Set("Content-Type", "application/json")

				h.MaxSumCalculator(c)

				Expect(c.Writer.Status()).To(Equal(http.StatusBadRequest))

			})
		})
	})
})

// StringReader, bir string'i io.Reader'a çeviren yardımcı bir fonksiyondur.
func StringReader(s string) *strings.Reader {
	return strings.NewReader(s)
}
