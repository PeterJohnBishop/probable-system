package main

import (
	// "probable-system/main.go/processing"

	"probable-system/main.go/server"
	"probable-system/main.go/server/handlers"
)

func main() {

	// Uncomment the following lines to generate data files, be sure to add output directory to gitignore

	// processing.GenerateTripData()
	// processing.GenerateRouteData()
	// processing.GenerateShapesData()
	// processing.GenerateStopTimesData()
	// processing.GenerateStopsData()

	go handlers.InitRouteMap()
	go handlers.InitShapesMap()
	go handlers.InitStopTimesMap()
	go handlers.InitStopsMap()
	go handlers.InitTripsMap()

	server.StartServer()

}
