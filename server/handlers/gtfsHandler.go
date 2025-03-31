package handlers

import (
	"fmt"
	"net/http"
	"probable-system/main.go/processing"
	"probable-system/main.go/processing/output"
	"probable-system/main.go/server/services/transportation"
)

var routeMap = make(map[string]processing.Route)

func initRouteMap() {
	for _, route := range output.Routes {
		routeMap[route.RouteID] = route
	}
}

func findRouteyID(routeId string) (processing.Route, bool) {
	fmt.Println("Finding route by ID: ", routeId)
	route, found := routeMap[routeId]
	if !found {
		fmt.Println("Route not found")
	} else {
		fmt.Println("Route found: ", route)

	}
	return route, found
}

func HandleGTFSRT(w http.ResponseWriter, r *http.Request) {

	initRouteMap()
	feed, err := transportation.FetchVehiclePosition()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching GTFS-RT: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	for _, entity := range feed.Entity {
		if entity.Vehicle != nil {

			route, _ := findRouteyID(*entity.Vehicle.Trip.RouteId)

			fmt.Fprintf(w, "Route: %s, Route: %s, Vehicle ID: %s, Lat: %f, Long: %f, Bearing: %f\n",
				route.RouteLongName,
				*entity.Vehicle.Trip.RouteId,
				*entity.Vehicle.Vehicle.Id,
				*entity.Vehicle.Position.Latitude,
				*entity.Vehicle.Position.Longitude,
				*entity.Vehicle.Position.Bearing)
		}
	}
}
