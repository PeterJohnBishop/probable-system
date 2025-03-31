package main

import (
	"fmt"
	"probable-system/main.go/processing"
	"sync"

	"probable-system/main.go/server"
	"probable-system/main.go/server/handlers"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		fmt.Println("Starting GenerateTripData...")
		haveData := processing.GenerateTripData()
		if haveData {
			fmt.Println("Initializing Trip Map...")
			handlers.InitTripsMap()
		}
		fmt.Println("Finished GenerateTripData")
		wg.Done()
	}()
	go func() {
		fmt.Println("Starting GenerateRouteData...")
		haveData := processing.GenerateRouteData()
		if haveData {
			fmt.Println("Initializing Route Map...")
			handlers.InitRouteMap()
		}
		fmt.Println("Finished GenerateRouteData")
		wg.Done()
	}()
	go func() {
		fmt.Println("Starting GenerateShapesData...")
		haveData := processing.GenerateShapesData()
		if haveData {
			fmt.Println("Initializing Shapes Map...")
			handlers.InitShapesMap()
		}
		fmt.Println("Finished GenerateShapesData")
		wg.Done()
	}()
	go func() {
		fmt.Println("Starting GenerateStopTimesData...")
		haveData := processing.GenerateStopTimesData()
		if haveData {
			fmt.Println("Initializing Stop Times Map...")
			handlers.InitStopTimesMap()
		}
		fmt.Println("Finished GenerateStopTimesData")
		wg.Done()
	}()
	go func() {
		fmt.Println("Starting GenerateStopsData...")
		haveData := processing.GenerateStopsData()
		if haveData {
			fmt.Println("Initializing Stops Map...")
			handlers.InitStopsMap()
		}
		fmt.Println("Finished GenerateStopsData")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("All processing tasks completed.")

	// Start the server
	server.StartServer()
}
