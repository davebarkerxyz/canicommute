package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-yaml/yaml"
	"googlemaps.github.io/maps"
)

type Config struct {
	ArrivalTime struct {
		Hour int `yaml:"hour"`
		Min  int `yaml:"min"`
	} `yaml:"arrival_time"`
	Locations  []string `yaml:"locations"`
	AutoSuffix string   `yaml:"auto_suffix"`
	ApiKey     string   `yaml:"api_key"`
}

func die(errText string, args ...any) {
	fmt.Fprintf(os.Stderr, errText+"\n", args...)
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		die(`Check commute times by public transit from given location to a list of configured destinations.

Usage: canicommute <location>

For example: canicommute "G1 1XH"`)
	}

	config := readConfig()
	location := os.Args[1]

	if len(config.AutoSuffix) != 0 {
		location += ", " + config.AutoSuffix
	}

	fmt.Printf("Getting commute time from \"%s\" to the following locations (arriving at %02d%02d):\n", location, config.ArrivalTime.Hour, config.ArrivalTime.Min)

	for _, dest := range config.Locations {
		fmt.Printf("- %s\n", dest)
	}
	fmt.Println("")

	matrixRows := getDistanceMatrix(config, location)
	printResults(matrixRows)
}

// Read config from config.yaml and return populated Config struct
func readConfig() Config {
	configFile, err := os.ReadFile("config.yaml")
	if err != nil {
		die("Error reading config file: %s", err)
	}

	config := Config{}
	err = yaml.UnmarshalStrict(configFile, &config)
	if err != nil {
		die("Error parsing config file: %s", err)
	}

	return config
}

// Call Google Maps Distance Matrix API for given config (origin, destinations, etc)
func getDistanceMatrix(config Config, location string) *maps.DistanceMatrixResponse {
	client, err := maps.NewClient(maps.WithAPIKey(config.ApiKey))
	if err != nil {
		die("Error connecting to Google Maps API: %s", err)
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
		die("Error making Google Maps Distance Matrix Request: %s", err)
	}

	return matrix
}

func printResults(matrix *maps.DistanceMatrixResponse) {
	fmt.Println("Results:")

	for destNum, row := range matrix.Rows[0].Elements {
		distance := row.Distance.HumanReadable
		if row.Distance.Meters == 0 {
			distance = "0m"
		}
		fmt.Printf("- %s: %s (%s)\n", matrix.DestinationAddresses[destNum], row.Duration, distance)
	}
}

// Get the next (or current) working day at the specified time
func getNextWorkingDay(hour int, min int) time.Time {
	now := time.Now()
	year, month, day := now.Date()
	nwd := time.Date(year, month, day, hour, min, 0, 0, time.Local)
	if now.Weekday() == time.Saturday {
		nwd = nwd.AddDate(0, 0, 2)
	} else if now.Weekday() == time.Sunday {
		nwd = nwd.AddDate(0, 0, 1)
	}

	return nwd
}
