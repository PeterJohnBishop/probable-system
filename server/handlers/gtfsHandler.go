package handlers

import (
	"fmt"
	"net/http"
	"probable-system/main.go/processing"
	"probable-system/main.go/processing/output"

	"probable-system/main.go/server/services/transportation"
)

var RoutesMap = make(map[string]processing.Route)
var ShapesMap = make(map[string]processing.Shape)
var StopTimesMap = make(map[string]processing.StopTime)
var StopsMap = make(map[string]processing.Stop)
var TripsMap = make(map[string]processing.Trip)

func InitRouteMap() {
	for _, route := range output.Routes {
		RoutesMap[route.RouteID] = route
	}
	fmt.Print("RoutesMap initialized with ", len(RoutesMap), " routes\n")
}
func InitShapesMap() {
	for _, shape := range output.Shapes {
		ShapesMap[shape.ShapeID] = shape
	}
	fmt.Print("ShapesMap initialized with ", len(ShapesMap), " shapes\n")
}
func InitStopTimesMap() {
	for _, stopTime := range output.StopTime {
		StopTimesMap[stopTime.TripID] = stopTime
	}
	fmt.Print("StopTimesMap initialized with ", len(StopTimesMap), " stop times\n")
}
func InitStopsMap() {
	for _, stop := range output.Stop {
		StopsMap[stop.StopID] = stop
	}
	fmt.Print("StopsMap initialized with ", len(StopsMap), " stops\n")
}
func InitTripsMap() {
	for _, trip := range output.Trips {
		TripsMap[trip.TripID] = trip
	}
	fmt.Print("TripsMap initialized with ", len(TripsMap), " trips\n")
}

func findRouteyID(routeId string) (processing.Route, bool) {
	route, found := RoutesMap[routeId]
	if !found {
		return processing.Route{}, false
	} else {
		return route, true

	}
}

func HandleGTFSRT(w http.ResponseWriter, r *http.Request) {

	feed, err := transportation.FetchVehiclePosition()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching GTFS-RT: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	for _, entity := range feed.Entity {
		if entity.Vehicle != nil {

			route, _ := findRouteyID(*entity.Vehicle.Trip.RouteId)

			var responseString string

			if entity.Vehicle == nil {
				responseString += "entity.Vehicle is nil\n"
			} else if entity.Vehicle.CurrentStatus == nil {
				responseString += "entity.Vehicle.CurrentStatus is nil\n"
			} else {
				responseString += fmt.Sprintf("CurrentStatus: %v\n", *entity.Vehicle.CurrentStatus)
			}
			responseString += fmt.Sprintf("RouteId: %s, Route: %s, Desc: %s, Type: %d, \n",
				*entity.Vehicle.Trip.RouteId,
				route.RouteLongName,
				route.RouteDesc,
				route.RouteType,
			)

			responseString += fmt.Sprintf("Vehicle ID: %s, Lat: %f, Long: %f, Bearing: %f\n",
				*entity.Vehicle.Vehicle.Id,
				*entity.Vehicle.Position.Latitude,
				*entity.Vehicle.Position.Longitude,
				*entity.Vehicle.Position.Bearing)

			fmt.Fprintf(w, "%s", responseString)
		}
	}
}
