package config

var (
	ServerName = ""
	ServerHost = "0.0.0.0"
	ServerPort = "8080"

	ServerDebug = true // TODO: Change this to false when deploying

	ServerSecret = "" // TODO: Change this to a more secure secret

	ServerLogToFile = true
	ServerLog       = "server.log"
	GinLog          = "gin.log"
)
