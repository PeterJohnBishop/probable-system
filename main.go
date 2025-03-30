package main

import (
	"probable-system/main.go/processing"
	"probable-system/main.go/server"
)

func main() {

	processing.GenerateTripData()

	server.StartServer()

}
