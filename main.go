package main

import (
	"flag"
	"fmt"
	"os"
	"stations/pkg/pathfinding"
	"stations/pkg/reader"
	"strconv"
)

func main() {
	// Parse command line arguments
	flag.Parse()

	// Ensure exactly 4 command line arguments are provided
	args := flag.Args()
	if len(args) != 4 {
		fmt.Fprintln(os.Stderr, "Error: Expected exactly 4 command line arguments")
		os.Exit(1)
	}

	// Extract arguments
	filename := args[0]
	startStation := args[1]
	endStation := args[2]
	numTrainsStr := args[3]

	// Validate the number of trains
	numTrains, err := strconv.Atoi(numTrainsStr)
	if err != nil || numTrains <= 0 {
		fmt.Fprintln(os.Stderr, "Error: Number of trains must be a positive integer")
		os.Exit(1)
	}

	// Load and parse the network map
	stations, connections, parseErrs := reader.ParseInputFile(filename)
	if len(parseErrs) > 0 {
		fmt.Fprintln(os.Stderr, "Error:")
		for _, parseErr := range parseErrs {
			fmt.Fprintln(os.Stderr, parseErr)
		}
		os.Exit(1)
	}

	// Check if the start and end stations are different and exist in the network
	if startStation == endStation {
		fmt.Fprintln(os.Stderr, "error: start and end station cannot be the same")
		os.Exit(1)
	}
	if _, exists := stations[startStation]; !exists {
		fmt.Fprintln(os.Stderr, "error: start station does not exist in the network")
		os.Exit(1)
	}
	if _, exists := stations[endStation]; !exists {
		fmt.Fprintln(os.Stderr, "error: end station does not exist in the network")
		os.Exit(1)
	}

	// Ensure a path exists between the start and end stations
	if !pathfinding.ThereIsPath(connections, startStation, endStation) {
		fmt.Fprintln(os.Stderr, "error: no connection exists between the start and end stations")
		os.Exit(1)
	}

	// Initialize stations and their edges
	for from, edges := range connections {
		stationFrom := stations[from]

		for _, to := range edges {
			distance := pathfinding.DummyWeight()
			stationTo := stations[to]
			stationFrom.AddEdge(stationTo, distance)
		}
	}

	// Find paths for the specified number of trains
	var trainPaths = pathfinding.Paths(stations[startStation], stations[endStation], numTrains, stations)

	// Simulate the movement of trains along the paths
	pathfinding.Simulate(numTrains, trainPaths, stations[startStation], stations[endStation])
}
