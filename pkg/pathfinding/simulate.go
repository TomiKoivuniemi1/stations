package pathfinding

import "fmt"

// Simulate the movement of multiple trains along their respective paths
func Simulate(numTrains int, trainPaths [][]*Station, start, end *Station) {
	// Create a slice of pointers to Train objects to hold the information of each train
	trains := make([]*Train, numTrains)

	// Initialize each train with a name, its path, and set its starting position to 0
	for i := 0; i < numTrains; i++ {
		trains[i] = &Train{Name: fmt.Sprintf("T%d", i+1), Path: trainPaths[i], Position: 0}
	}

	// Map to keep track of which stations are currently occupied by which trains
	stationOccupancy := make(map[*Station]*Train)
	// Boolean to check if a direct connection exists for the current train
	directConnection := false
	// Boolean to track if a direct connection has been reserved
	directConnectionReserved := false

	// Infinite loop to simulate the movement of the trains
	for {
		// Flag to track if all trains have completed their paths
		allTrainsCompleted := true

		// Iterate through each train to update its position
		for _, train := range trains {
			// Skip this train if it has already completed its journey
			if train.Completed {
				continue
			}
			// If at least one train is still in motion, set the flag to false
			allTrainsCompleted = false

			// Check if the train's path includes a direct connection to the end station
			directConnection = CheckDirect(train.Path, end)
			// Get the current station based on the train's position in its path
			currentStation := train.Path[train.Position]

			// Check if the train has not yet reached the end of its path
			if train.Position < len(train.Path)-1 {
				// Determine the next station in the train's path
				nextStation := train.Path[train.Position+1]

				// Check if the next station is unoccupied or is the start or end station
				if stationOccupancy[nextStation] == nil || nextStation == start || nextStation == end {
					// Free up the current station in stationOccupancy if it's not the start or end station
					if currentStation != start && currentStation != end {
						stationOccupancy[currentStation] = nil
					}

					// If the train is on a direct connection and it's reserved, the train must wait
					if directConnection && directConnectionReserved {
						continue
					} else if directConnection && !directConnectionReserved {
						// Reserve the direct connection for this train
						directConnectionReserved = true
					}

					// Move the train to the next station by incrementing its position
					train.Position++
					// Mark the next station as occupied by this train if it's not the end station
					if nextStation != end {
						stationOccupancy[nextStation] = train
					}
				}
			} else {
				// If the train has reached the end of its path, mark it as completed
				train.Completed = true
				// If it's on a direct connection, free up the reservation
				if directConnection {
					directConnectionReserved = false
				}
			}
		}

		// Flag to check if any train has moved during this iteration
		moved := false
		// Print the current position of each train that has moved
		for _, train := range trains {
			if !train.Completed && train.Position != 0 {
				// Print the train's name and its current station
				fmt.Printf("%s-%s ", train.Name, train.Path[train.Position].Name)
				moved = true
			}
		}
		// If any train moved, print a newline for formatting
		if moved {
			fmt.Println()
		}

		// If all trains have completed their paths, exit the loop
		if allTrainsCompleted {
			break
		}
	}
}
