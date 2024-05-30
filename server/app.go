package server

import (
	"fmt"
	"log"
	"os"

	"api-server/config"
	"github.com/gin-gonic/gin"
)

func init() {
	config.ServerName = os.Getenv("SERVER_DOMAIN")
	config.ServerSecret = os.Getenv("SERVER_SECRET_KEY")

	if os.Getenv("SERVER_IS_DEV") == "false" {
		config.ServerDebug = false
	} else {
		config.ServerDebug = true
	}

	if os.Getenv("SERVER_LOG_TO_FILE") == "false" {
		config.ServerLogToFile = false
	} else {
		config.ServerLogToFile = true
	}
}

func Start() error {
	if config.ServerLogToFile {
		serverLogFile, err := os.OpenFile(config.ServerLog, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}

		ginLogFile, err := os.OpenFile(config.GinLog, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}

		log.SetOutput(serverLogFile)
		gin.DefaultWriter = ginLogFile
	}

	if config.ServerDebug {
		gin.SetMode(gin.DebugMode)

		log.Println("Running in debug mode")
	} else {
		gin.SetMode(gin.ReleaseMode)

		log.Println("Running in release mode")
	}

	r := NewRouter()

	addr := fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort)
	log.Println("Starting server on", addr)

	return r.Run(addr)
}
