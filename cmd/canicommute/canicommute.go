package main

import (
	"fmt"
	"os"

	"github.com/davebarkerxyz/canicommute/commute"
)

func main() {
	if len(os.Args) != 2 {
		commute.Die(`Check commute times by public transit from given location to a list of configured destinations.

Usage: canicommute <location>

For example: canicommute "G1 1XH"`)
	}

	config := commute.GetConfig()
	location := os.Args[1]

	if len(config.AutoSuffix) != 0 {
		location += ", " + config.AutoSuffix
	}

	fmt.Printf("Getting commute time from \"%s\" to the following locations (arriving at %02d%02d):\n", location, config.ArrivalTime.Hour, config.ArrivalTime.Min)

	for _, dest := range config.Locations {
		fmt.Printf("- %s\n", dest)
	}
	fmt.Println("")

	matrixRows := commute.GetDistanceMatrix(config, location)
	commute.PrintResults(matrixRows)
}
