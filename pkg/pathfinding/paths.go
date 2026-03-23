package pathfinding

// Paths finds the fastest paths for multiple trains from the start station to the end station
func Paths(startStation *Station, endStation *Station, numTrains int, stations map[string]*Station) [][]*Station {
	// Map to keep track of edges that should be excluded for each train's path
	excludedEdges := make(map[*Edge]bool)
	trainPaths := make([][]*Station, numTrains)

	for i := 0; i < numTrains; i++ {
		// Reset the excludedEdges map if the number of excluded edges is a multiple of the number of edges from the starting station
		if len(excludedEdges)%len(startStation.Edges) == 0 {
			excludedEdges = make(map[*Edge]bool)
		}

		// The last train takes the fastest path with no edge exclusions
		if i == numTrains-1 {
			excludedEdges = make(map[*Edge]bool) // Ensure no edges are excluded for the last train
			trainPaths[i] = SinglePath(startStation, endStation, excludedEdges, trainPaths[:i], stations)
		} else {
			// For other trains, find the fastest path considering previously excluded edges and paths of other trains
			trainPaths[i] = SinglePath(startStation, endStation, excludedEdges, trainPaths[:i], stations)
		}
	}

	return trainPaths
}
