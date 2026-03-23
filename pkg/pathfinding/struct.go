package pathfinding

// Station represents a station in the pathfinding problem.
type Station struct {
	Name  string   // Name of the station.
	X     int      // X-coordinate of the station.
	Y     int      // Y-coordinate of the station.
	Dist  float64  // Distance from the source station (used in pathfinding algorithms).
	Prev  *Station // Previous station in the path (used for path reconstruction).
	Edges []*Edge  // List of edges (connections) from this station to other stations.
}

// Train represents a train in the pathfinding problem.
type Train struct {
	Name      string     // Name of the train.
	Path      []*Station // Path that the train follows.
	Position  int        // Index of the current station in the Path slice.
	Completed bool       // Flag indicating if the train has completed its journey.
}

// Connection represents a connection between two stations.
type Connection struct {
	Start, End string // Names of the start and end stations for the connection.
}

// Edge represents a connection between two stations with a specific weight.
type Edge struct {
	From, To *Station // The stations connected by this edge.
	Weight   float64  // The weight (distance or cost) of this edge.
}

// QueuePriority implements a priority queue for stations.
// The priority queue is used to manage stations in algorithms such as Dijkstra's.
type QueuePriority []*Station

// Len returns the number of elements in the priority queue
func (pq QueuePriority) Len() int { return len(pq) }

// Less compares two stations based on their distance, used to order the queue
func (pq QueuePriority) Less(i, j int) bool {
	return pq[i].Dist < pq[j].Dist
}

// Swap exchanges the positions of two stations in the priority queue
func (pq QueuePriority) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push adds a new station to the priority queue
func (pq *QueuePriority) Push(x any) {
	item := x.(*Station)
	*pq = append(*pq, item)
}

// Pop removes and returns the station with the smallest distance from the priority queue
func (pq *QueuePriority) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}
