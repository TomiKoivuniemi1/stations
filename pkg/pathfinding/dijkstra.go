package pathfinding

import (
	"container/heap"
	"math"
)

// Dijkstra finds the shortest path from startStation to endStation,
// considering excluded edges and additional costs due to other train paths.
func Dijkstra(startStation *Station, endStation *Station, excludedEdges map[*Edge]bool, otherTrainPaths [][]*Station) ([]*Station, float64) {
	// Initialize a priority queue to manage stations to be processed, sorted by distance.
	pq := make(QueuePriority, 0)
	heap.Init(&pq)
	heap.Push(&pq, startStation) // Add the start station to the priority queue.
	startStation.Dist = 0        // Distance from the start station to itself is zero.

	// Map to keep track of the previous station for each station in the path.
	prev := make(map[*Station]*Station)

	// Process stations until the priority queue is empty.
	for pq.Len() > 0 {
		station := heap.Pop(&pq).(*Station) // Get the station with the smallest distance from the queue.

		// If the distance to this station is infinity, it means it's unreachable.
		if station.Dist == math.Inf(1) {
			continue
		}

		// Check all edges from the current station.
		for _, edge := range station.Edges {
			// Skip edges that are excluded.
			if excludedEdges[edge] {
				continue
			}

			// Calculate additional cost for using this edge due to other trains' paths.
			additionalCost := EdgeCost(edge, otherTrainPaths)
			// Calculate the tentative distance to the destination station of this edge.
			tentativeDist := station.Dist + edge.Weight + additionalCost

			// Update the distance and path if a shorter path is found.
			if tentativeDist < edge.To.Dist {
				edge.To.Dist = tentativeDist
				prev[edge.To] = station
				heap.Push(&pq, edge.To) // Add the destination station to the queue.
			}
		}
	}

	// Reconstruct the shortest path from endStation to startStation using the prev map.
	path := []*Station{}
	current := endStation
	for current != nil {
		path = append(path, current)
		current = prev[current]
	}

	// Reverse the path to be in the order from startStation to endStation.
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	// Mark the first edge in the path as excluded for future path calculations.
	for i := 0; i < 1; i++ {
		for _, edge := range path[i].Edges {
			if edge.To == path[i+1] {
				excludedEdges[edge] = true
				break
			}
		}
	}

	// Return the shortest path and the distance to the end station.
	return path, endStation.Dist
}
