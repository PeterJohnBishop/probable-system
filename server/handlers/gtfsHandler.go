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

func findRouteByID(routeId string) (processing.Route, bool) {
	route, found := RoutesMap[routeId]
	if !found {
		return processing.Route{}, false
	} else {
		return route, true

	}
}

func InitShapesMap() {
	for _, shape := range output.Shapes {
		ShapesMap[shape.ShapeID] = shape
	}
	fmt.Print("ShapesMap initialized with ", len(ShapesMap), " shapes\n")
}

func findShapeById(shapeId string) (processing.Shape, bool) {
	shape, found := ShapesMap[shapeId]
	if !found {
		return processing.Shape{}, false
	} else {
		return shape, true
	}

}

func InitStopTimesMap() {
	for _, stopTime := range output.StopTime {
		StopTimesMap[stopTime.TripID] = stopTime
	}
	fmt.Print("StopTimesMap initialized with ", len(StopTimesMap), " stop times\n")
}

func findStopTimeById(tripId string) (processing.StopTime, bool) {
	stopTime, found := StopTimesMap[tripId]
	if !found {
		return processing.StopTime{}, false
	} else {
		return stopTime, true
	}
}

func InitStopsMap() {
	for _, stop := range output.Stop {
		StopsMap[stop.StopID] = stop
	}
	fmt.Print("StopsMap initialized with ", len(StopsMap), " stops\n")
}

func findStopById(stopId string) (processing.Stop, bool) {
	stop, found := StopsMap[stopId]
	if !found {
		return processing.Stop{}, false
	} else {
		return stop, true
	}
}

func InitTripsMap() {
	for _, trip := range output.Trips {
		TripsMap[trip.TripID] = trip
	}
	fmt.Print("TripsMap initialized with ", len(TripsMap), " trips\n")
}

func findTripById(tripId string) (processing.Trip, bool) {
	trip, found := TripsMap[tripId]
	if !found {
		return processing.Trip{}, false
	} else {
		return trip, true
	}
}

func HandleAlert(w http.ResponseWriter, r *http.Request) {
	feed, err := transportation.FetchAlerts()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching GTFS-RT: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	for _, entity := range feed.Entity {
		if entity.Alert != nil {
			var responseString string
			responseString += fmt.Sprintf("Response: %v\n", entity)
			fmt.Fprintf(w, "%s", responseString)
		}
	}

	// Response: id: "49932" alert: {active_period: {start: 1725542640 end: 1748595540
	// 	} informed_entity: {agency_id: "RTD" route_id: "11" route_type: 3 stop_id: "15398"
	// 	} informed_entity: {agency_id: "RTD" route_id: "11" route_type: 3 stop_id: "15411"
	// 	} informed_entity: {agency_id: "RTD" route_id: "11" route_type: 3 stop_id: "15417"
	// 	} informed_entity: {agency_id: "RTD" route_id: "11" route_type: 3 stop_id: "15397"
	// 	} informed_entity: {agency_id: "RTD" route_id: "11" route_type: 3 stop_id: "15410"
	// 	} cause:CONSTRUCTION effect:DETOUR header_text: {translation: {text: "Route 11 detoured from 7:24 AM today through Thu May 29 due to construction." language: "en"
	// 		}
	// 	} description_text: {translation: {text: "Affected stops:\r\nMississippi Ave & S Leyden St (#15397) (eastbound)\r\nMississippi Ave & S Monaco Pkwy (#15410) (eastbound)\r\nMississippi Ave & S Leyden St (#15398) (westbound)\r\nMississippi Ave & S Monaco Pkwy (#15411) (westbound)\r\nMississippi Ave & S Oneida St (#15417) (westbound)\r\n\nFor eastbound Route 11 get on/off buses at:\r\nTennessee Ave & S Jersey St (#16741)\r\nMississippi Ave & S Oneida St (#15416)\r\n\nFor westbound Route 11 get on/off buses at:\r\nMississippi Ave & S Quebec St (#15428)\r\nTennessee Ave & S Jersey St (#16742)" language: "en"
	// 		}
	// 	}
	// }
}

func HandleTripUpdate(w http.ResponseWriter, r *http.Request) {
	feed, err := transportation.FetchTripUpdates()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching GTFS-RT: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	for _, entity := range feed.Entity {
		if entity.TripUpdate != nil {
			var responseString string
			responseString += fmt.Sprintf("Response: %v\n", entity)
			fmt.Fprintf(w, "%s", responseString)
		}
	}

	//  Response: id: "1743462284_115184047" trip_update: {trip: {trip_id: "115184047" route_id: "0" direction_id: 0 schedule_relationship:SCHEDULED
	// 	} vehicle: {id: "3176E117552F580BE063DC4D1FAC4CA4" label: "9329"
	// 	} stop_time_update: {stop_sequence: 35 stop_id: "17834" arrival: {time: 1743462056
	// 		} departure: {time: 1743462056
	// 		} schedule_relationship:SCHEDULED
	// 	} stop_time_update: {stop_sequence: 36 stop_id: "34343" arrival: {time: 1743462146
	// 		} departure: {time: 1743462146
	// 		} schedule_relationship:SCHEDULED
	// 	} timestamp: 1743462269
	//  }
}

func HandleVehiclePosition(w http.ResponseWriter, r *http.Request) {

	feed, err := transportation.FetchVehiclePosition()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching GTFS-RT: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	for _, entity := range feed.Entity {

		trip, foundTrip := findTripById(*entity.Vehicle.Trip.TripId)
		route, foundRoute := findRouteByID(*entity.Vehicle.Trip.RouteId)
		stop, foundStop := findStopById(*entity.Vehicle.StopId)
		if entity.Vehicle != nil {
			var responseString string
			if foundTrip {
				responseString += fmt.Sprintf("Trip: %v\n", trip)
			} else {
				responseString += fmt.Sprintf("Trip not found for Trip ID: %s\n", *entity.Vehicle.Trip.TripId)
			}
			if foundRoute {
				responseString += fmt.Sprintf("Route: %v\n", route)
			} else {
				responseString += fmt.Sprintf("Route not found for Route ID: %s\n", *entity.Vehicle.Trip.RouteId)
			}
			responseString += fmt.Sprintf("direction_id: %d\n", *entity.Vehicle.Trip.DirectionId)
			responseString += fmt.Sprintf("schedule_relationship: %s\n", *entity.Vehicle.Trip.ScheduleRelationship)
			responseString += fmt.Sprintf("vehicle_id: %s\n", *entity.Vehicle.Vehicle.Id)
			responseString += fmt.Sprintf("vehicle_label: %s\n", *entity.Vehicle.Vehicle.Label)
			responseString += fmt.Sprintf("latitude: %f\n", *entity.Vehicle.Position.Latitude)
			responseString += fmt.Sprintf("longitude: %f\n", *entity.Vehicle.Position.Longitude)
			responseString += fmt.Sprintf("bearing: %f\n", *entity.Vehicle.Position.Bearing)
			if foundStop {
				responseString += fmt.Sprintf("Stop: %v\n", stop)
			} else {
				responseString += fmt.Sprintf("Stop not found for Stop ID: %s\n", *entity.Vehicle.StopId)
			}
			responseString += fmt.Sprintf("current_status: %s\n", *entity.Vehicle.CurrentStatus)
			responseString += fmt.Sprintf("timestamp: %d\n", *entity.Vehicle.Timestamp)
			responseString += fmt.Sprintf("occupancy_status: %s\n", *entity.Vehicle.OccupancyStatus)
			fmt.Fprintf(w, "%s", responseString)
		}
	}

	// Response: id: "1743462390_3176E117547F580BE063DC4D1FAC4CA4" vehicle: {trip: {trip_id: "115198391" route_id: "520" direction_id: 0 schedule_relationship:SCHEDULED
	// } vehicle: {id: "3176E117547F580BE063DC4D1FAC4CA4" label: "6548"
	// } position: {latitude: 39.9867 longitude: -104.805916 bearing: 92
	// } stop_id: "20033" current_status:IN_TRANSIT_TO timestamp: 1743462366 occupancy_status:EMPTY
	// }
}
