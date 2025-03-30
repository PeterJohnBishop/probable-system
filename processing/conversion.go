package processing

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func GenerateTripData() {
	file, err := os.Open("/Users/peterbishop/Development/probable-system/processing/input/trips.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
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

	outputFile := "/Users/peterbishop/Development/probable-system/processing/output/trips.json"
	err = os.WriteFile(outputFile, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}

	fmt.Println("JSON data successfully saved to", outputFile)
}
