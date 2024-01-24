package main

import (
	"fmt"

	"github.com/acergiyen/qlub-golang-app-challange/internal/app/config"
	"github.com/acergiyen/qlub-golang-app-challange/internal/app/handler"
	"github.com/acergiyen/qlub-golang-app-challange/internal/app/prepare"
	"github.com/acergiyen/qlub-golang-app-challange/internal/usecases/maxsumcalculator"
	"github.com/gin-gonic/gin"
)

func main() {
	// Get configuration settings
	config, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	// Set up logger
	logger := prepare.AppLogger(config)
	logger.Printf("%v application started", config.App.Name)

	// Initialize max sum calculator
	calculator := maxsumcalculator.NewMaxSumCalculator(logger)

	// Initialize request handler
	handler := handler.NewHandler(calculator)

	// Set up Gin router
	r := gin.Default()
	r.POST("/calculateMaxSum", handler.MaxSumCalculator)

	// Start the HTTP server
	logger.Printf("Server is running on %v", config.App.Port)
	address := fmt.Sprintf(":%d", config.App.Port)
	r.Run(address)
}
