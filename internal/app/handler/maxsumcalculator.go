package handler

import (
	"net/http"

	"github.com/acergiyen/qlub-golang-app-challange/internal/usecases/maxsumcalculator"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	calculator *maxsumcalculator.MaxSumCalculator
}

func NewHandler(calculator *maxsumcalculator.MaxSumCalculator) *Handler {
	return &Handler{calculator: calculator}
}

func (h *Handler) MaxSumCalculator(c *gin.Context) {
	var treeRequest maxsumcalculator.TreeRequest
	if err := c.ShouldBindJSON(&treeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
		return
	}

	maxPathSum := h.calculator.Handle(treeRequest)

	c.JSON(http.StatusOK, gin.H{"maxPathSum": maxPathSum})

}
