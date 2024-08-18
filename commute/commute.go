package commute

import (
	"context"
	"fmt"
	"strconv"

	"googlemaps.github.io/maps"
)

// Call Google Maps Distance Matrix API for given config (origin, destinations, etc)
func GetDistanceMatrix(config Config, location string) *maps.DistanceMatrixResponse {
	client, err := maps.NewClient(maps.WithAPIKey(config.ApiKey))
	if err != nil {
		Die("Error connecting to Google Maps API: %s", err)
	}

	arrivalTime := getNextWorkingDay(config.ArrivalTime.Hour, config.ArrivalTime.Min)

	req := maps.DistanceMatrixRequest{
		Origins:      []string{location},
		Destinations: config.Locations,
		ArrivalTime:  strconv.FormatInt(arrivalTime.Unix(), 10),
		Mode:         maps.TravelModeTransit,
	}

	matrix, err := client.DistanceMatrix(context.Background(), &req)
	if err != nil {
		Die("Error making Google Maps Distance Matrix Request: %s", err)
	}

	return matrix
}

func PrintResults(matrix *maps.DistanceMatrixResponse) {
	fmt.Println("Results:")

	for destNum, row := range matrix.Rows[0].Elements {
		distance := row.Distance.HumanReadable
		if row.Distance.Meters == 0 {
			distance = "0m"
		}
		fmt.Printf("- %s: %s (%s)\n", matrix.DestinationAddresses[destNum], row.Duration, distance)
	}
}
