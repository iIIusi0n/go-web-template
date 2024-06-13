package main

import (
	"api-server/server"
)

func main() {
	err := server.Start()
	if err != nil {
		panic(err)
	}
}
