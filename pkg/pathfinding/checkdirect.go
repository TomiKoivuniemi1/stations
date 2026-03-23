package pathfinding

// checkDirect determines if the provided path represents a direct connection to the end station.
// A direct connection is defined as a path with exactly two stations where the last station is the end station.
func CheckDirect(path []*Station, end *Station) bool {
	// Check if the path consists of exactly two stations and the last station is the end station.
	return len(path) == 2 && path[len(path)-1] == end
}
