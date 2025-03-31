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
	fmt.Println("Finding route by ID: ", routeId)
	route, found := RoutesMap[routeId]
	if !found {
		fmt.Println("Route not found")
	} else {
		fmt.Println("Route found: ", route)

	}
	return route, found
}

func currentStatus(status int) string {
	switch status {
	case 0:
		return "INCOMING_AT"
	case 1:
		return "STOPPED_AT"
	case 2:
		return "IN_TRANSIT_TO"
	default:
		return "UNKNOWN"
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

			// vehicle
			// 	trip
			// 		trip_id
			// 		route_id
			// 		direction_id
			// 		schedule_relationship. SCHEDULED if trip is running as scheduled, ADDED if trip is an added trip, or CANCELED if trip has been canceled.
			// 	vehicle
			// 		id
			// 		label
			// 		position
			// 			latitude
			// 			longitude
			// 			bearing
			// 	stop_id
			// 	current_status. Will have either a 0, 1, or 2.
			// 	0 = INCOMING_AT
			// 	1 = STOPPED_AT - if vehicle is stopped at the stop_id
			// 	2 = IN_TRANSIT_TO - if vehicle is on its way to the stop_id
			// 	timestamp := entity.Vehicle.Trip.TripId

			var responseString string

			responseString += fmt.Sprintf("status: %s\n", currentStatus(int(*entity.Vehicle.CurrentStatus)))

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
