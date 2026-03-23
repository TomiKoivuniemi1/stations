package pathfinding

// EdgeCost calculates the additional cost of using a particular edge
// based on the paths of other trains. The cost increases if the edge
// is frequently used by other trains.
func EdgeCost(consideredEdge *Edge, otherTrainPaths [][]*Station) float64 {
	additionalCost := 0.0

	// Iterate over each path of other trains.
	for _, otherPath := range otherTrainPaths {
		// Traverse each segment of the train's path.
		for i := 0; i < len(otherPath)-1; i++ {
			// Check if the current segment matches the edge being considered.
			if otherPath[i] == consideredEdge.To {
				// Increment the additional cost if the edge is found in the path.
				additionalCost += 1.0
			}
		}
	}

	// Return the total additional cost for the edge.
	return additionalCost
}
