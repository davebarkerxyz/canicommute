package main

import (
	"fmt"
	"os"

	"github.com/davebarkerxyz/canicommute/commute"
)

func main() {
	if len(os.Args) == 1 {
		printUsage()
	}

	switch os.Args[1] {
	case "check":
		if len(os.Args) != 3 {
			printUsage()
		}
		checkCommute(os.Args[2])
	case "serve":
		serve()
	default:
		printUsage()
	}
}

func printUsage() {
	commute.Die(`Check commute times by public transit from given location to a list of configured destinations.

Usage: canicommute
         check <location>    check commute time to location
         serve               start web server

For example: canicommute check "G1 1XH"`)
}

func checkCommute(location string) {
	config := commute.GetConfig()

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

func serve() {
	commute.Die("TODO: web server")
}
