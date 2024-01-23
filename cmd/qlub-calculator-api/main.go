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
	config, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	logger := prepare.AppLogger(config)
	logger.Printf("%v application started", config.App.Name)
	calculator := maxsumcalculator.NewMaxSumCalculator(logger)
	handler := handler.NewHandler(calculator)
	r := gin.Default()
	r.POST("/calculateMaxSum", handler.MaxSumCalculator)

	logger.Printf("Server is running on %v", config.App.Port)
	address := fmt.Sprintf(":%d", config.App.Port)
	r.Run(address)
}
