package handlers

import (
	"fmt"
	"net/http"
	"probable-system/main.go/server/services/transportation"
)

func HandleGTFSRT(w http.ResponseWriter, r *http.Request) {
	feed, err := transportation.FetchVehiclePosition()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching GTFS-RT: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	for _, entity := range feed.Entity {
		if entity.Vehicle != nil {
			fmt.Fprintf(w, "Vehicle ID: %s, Lat: %f, Long: %f Bearing: %f\n",
				*entity.Vehicle.Vehicle.Id,
				*entity.Vehicle.Position.Latitude,
				*entity.Vehicle.Position.Longitude,
				*entity.Vehicle.Position.Bearing)
		}
	}
}
