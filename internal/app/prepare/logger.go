package prepare

import (
	"fmt"
	"log"
	"os"

	"github.com/acergiyen/qlub-golang-app-challange/internal/app/config"
)

// AppLogger creates and returns a new logger with specific configurations.
func AppLogger(config *config.Config) *log.Logger {
	logger := log.New(os.Stdout, fmt.Sprintf("["+config.App.Name+"] "), log.Ldate|log.Ltime)

	return logger
}
