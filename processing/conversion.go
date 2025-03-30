package processing

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const outputUrl = "/Users/peterbishop/Development/probable-system/processing/output/"
const inputUrl = "/Users/peterbishop/Development/probable-system/processing/input/"

func OpenFile(fileName string) ([][]string, error) {
	file, err := os.Open(inputUrl + fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return nil, err
	}

	return records, nil
}

func GenerateTripData() {

	records, err := OpenFile("trips.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	var trips []Trip
	for i, row := range records {
		if i == 0 {
			continue
		}

		var directionID int
		fmt.Sscanf(row[4], "%d", &directionID)
		blockID := strings.TrimSpace(row[5])

		trips = append(trips, Trip{
			RouteID:      row[0],
			ServiceID:    row[1],
			TripID:       row[2],
			TripHeadsign: row[3],
			DirectionID:  directionID,
			BlockID:      blockID,
			ShapeID:      row[6],
		})
	}

	jsonData, err := json.MarshalIndent(trips, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	outputFile := fmt.Sprintf(outputUrl + "trip_data.json")
	err = os.WriteFile(outputFile, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}

	fmt.Println("JSON data successfully saved to", outputFile)
}

func GenerateRouteData() {
	records, err := OpenFile("routes.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	var routes []Route
	for i, row := range records {
		if i == 0 {
			continue
		}

		routeType := strings.TrimSpace(row[5])
		routeTypeInt := 0
		if routeType == "3" {
			routeTypeInt = 3
		}
		routes = append(routes, Route{
			RouteID:        row[0],
			AgencyID:       row[1],
			RouteShortName: row[2],
			RouteLongName:  row[3],
			RouteDesc:      row[4],
			RouteType:      routeTypeInt,
			RouteURL:       row[6],
			RouteColor:     row[7],
			RouteTextColor: row[8],
		})

	}

	jsonData, err := json.MarshalIndent(routes, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}
	outputFile := fmt.Sprintf(outputUrl + "route_data.json")
	err = os.WriteFile(outputFile, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}
	fmt.Println("JSON data successfully saved to", outputFile)
}

func GenerateShapesData() {
	records, err := OpenFile("shapes.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	var shapes []Shape
	for i, row := range records {
		if i == 0 {
			continue
		}

		shapeDistTraveled := strings.TrimSpace(row[4])
		shapeDistTraveledFloat := 0.0
		if shapeDistTraveled != "" {
			fmt.Sscanf(shapeDistTraveled, "%f", &shapeDistTraveledFloat)
		}
		shapePtSequence := strings.TrimSpace(row[3])
		shapePtSequenceInt := 0
		if shapePtSequence != "" {
			fmt.Sscanf(shapePtSequence, "%d", &shapePtSequenceInt)
		}
		shapes = append(shapes, Shape{
			ShapeID: row[0],
			ShapePtLat: func() float64 {
				lat, _ := strconv.ParseFloat(row[1], 64)
				return lat
			}(),
			ShapePtLon: func() float64 {
				lon, _ := strconv.ParseFloat(row[2], 64)
				return lon
			}(),
			ShapePtSequence:   shapePtSequenceInt,
			ShapeDistTraveled: shapeDistTraveledFloat,
		})

	}

	jsonData, err := json.MarshalIndent(shapes, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}
	outputFile := fmt.Sprint(outputUrl + "shape_data.json")
	err = os.WriteFile(outputFile, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}
	fmt.Println("JSON data successfully saved to", outputFile)
}

func GenerateStopTimesData() {
	records, err := OpenFile("stop_times.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	var stopTimes []StopTime
	for i, row := range records {
		if i == 0 {
			continue
		}
		stopSequence, _ := strconv.Atoi(row[4])
		pickupType, _ := strconv.Atoi(row[6])
		dropOffType, _ := strconv.Atoi(row[7])
		stopTimes = append(stopTimes, StopTime{
			TripID:        row[0],
			ArrivalTime:   row[1],
			DepartureTime: row[2],
			StopID:        row[3],
			StopSequence:  stopSequence,
			PickupType:    pickupType,
			DropOffType:   dropOffType,
		})
	}

	jsonData, err := json.MarshalIndent(stopTimes, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}
	outputFile := fmt.Sprint(outputUrl + "stop_time_data.json")
	err = os.WriteFile(outputFile, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}
	fmt.Println("StopTimes JSON successfully saved to", outputFile)
}

func GenerateStopsData() {
	records, err := OpenFile("stops.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	var stops []Stop
	for i, row := range records {
		if i == 0 {
			continue
		}
		lat, _ := strconv.ParseFloat(row[4], 64)
		lon, _ := strconv.ParseFloat(row[5], 64)
		stops = append(stops, Stop{
			StopID:   row[0],
			StopCode: row[1],
			StopName: row[2],
			StopDesc: row[3],
			StopLat:  lat,
			StopLon:  lon,
		})
	}

	jsonData, err := json.MarshalIndent(stops, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}
	outputFile := fmt.Sprint(outputUrl + "stop_data.json")
	err = os.WriteFile(outputFile, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}
	fmt.Println("Stops JSON successfully saved to", outputFile)
}
