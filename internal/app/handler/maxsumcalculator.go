package handler

import (
	"net/http"

	"github.com/acergiyen/qlub-golang-app-challange/internal/usecases/maxsumcalculator"
	"github.com/gin-gonic/gin"
)

// Handler handles HTTP requests related to the maximum path sum calculator.
type Handler struct {
	calculator *maxsumcalculator.MaxSumCalculator
}

// NewHandler creates a new instance of Handler with the given maximum path sum calculator.
func NewHandler(calculator *maxsumcalculator.MaxSumCalculator) *Handler {
	return &Handler{calculator: calculator}
}

// MaxSumCalculator handles the HTTP request for calculating the maximum path sum.
func (h *Handler) MaxSumCalculator(c *gin.Context) {
	// Parse the JSON request into the TreeRequest struct.
	var treeRequest maxsumcalculator.TreeRequest
	if err := c.ShouldBindJSON(&treeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
		return
	}

	// Call the Handle method of the MaxSumCalculator to calculate the maximum path sum.
	maxPathSum := h.calculator.Handle(treeRequest)

	// Respond with the calculated maximum path sum.
	c.JSON(http.StatusOK, gin.H{"maxPathSum": maxPathSum})
}
