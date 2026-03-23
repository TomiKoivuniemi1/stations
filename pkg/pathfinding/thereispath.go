package pathfinding

// ThereIsPath checks if there is a path between the start and end stations
// using Breadth-First Search (BFS). It does not perform any simulations,
// but simply verifies connectivity between the two stations.
func ThereIsPath(connections map[string][]string, start, end string) bool {
	visited := make(map[string]bool)
	queue := []string{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			return true
		}

		for _, neighbor := range connections[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	return false
}
