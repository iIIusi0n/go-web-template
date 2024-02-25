package main

import (
	"math/rand"
	"time"

	"api-server/server"
)

func main() {
	rand.Seed(time.Now().Unix())

	err := server.Start()
	if err != nil {
		panic(err)
	}
}
