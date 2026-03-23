package pathfinding

import (
	"math"
)

// Weight calculates the Euclidean distance between two stations based on their coordinates.
func Weight(from, to string, stations map[string]*Station) float64 {
	f := stations[from]
	t := stations[to]
	dx := float64(f.X - t.X)
	dy := float64(f.Y - t.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

// Method to add edges to each station
func (n *Station) AddEdge(to *Station, weight float64) {
	n.Edges = append(n.Edges, &Edge{To: to, Weight: weight})
}

func DummyWeight() float64 {
	return 1.0
}
