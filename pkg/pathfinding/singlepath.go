package pathfinding

import "math"

// SinglePath finds the shortest path from startStation to endStation considering excluded edges and paths of other trains.
// It initializes distances and previous stations, then uses Dijkstra's algorithm to find the path.
func SinglePath(startStation *Station, endStation *Station, excludedEdges map[*Edge]bool, otherTrainPaths [][]*Station, stations map[string]*Station) []*Station {
	// Initialize each station's distance to infinity and previous station to nil
	for _, station := range stations {
		station.Dist = math.Inf(1) // Set distance to positive infinity
		station.Prev = nil         // No previous station initially
	}
	startStation.Dist = 0 // Distance from the start station to itself is zero

	// Find the shortest path using Dijkstra's algorithm
	path, _ := Dijkstra(startStation, endStation, excludedEdges, otherTrainPaths)

	return path
}
