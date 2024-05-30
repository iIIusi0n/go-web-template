package config

const (
	ImageDirectory = "/usr/local/share/images"
)

var (
	MysqlHostname = "db"
	MysqlPort     = "3306"
	MysqlUsername = "" // TODO: Change this to a more readable username
	MysqlPassword = "" // TODO: Change this to a more secure password
	MysqlDatabase = "" // TODO: Change this to a more readable database name
)

const (
	MysqlDateTimeLayout = "2006-01-02 15:04:05"
)
