STATIONS

Welcome to stations-pathfinder!


Usage: go run . [path to file containing network map] [start station] [end station] [number of trains]

For example go run . network.map beethoven part 9

You may modify the network.map or create new maps for your needs.

Reviewers: You can use the makefile for testing. Just type "make" in terminal on the root folder. This will create a "testresults.txt" which will contain the results from premade commands of the review page tests starting from the top of the review page.

OR 

testing.txt includes commands in the same order as in the review page. Time to time the filenames, station names etc also hint where you are going if you get lost. There is no added numbering or hints in the commands so that it is easy to just copy paste to terminal for testing. You can also paste multiple commands at the same time and it will give the results in given order.

Code Execution Flow:
1. Loading the Map: The program starts by loading the map file and parsing the stations and connections.
2. Pathfinding: The program calculates paths for each train using Dijkstra’s algorithm, adjusting for previously occupied paths.
3. Simulation: Once paths are calculated, the simulation of train movements begins, considering station occupancy and path conflicts.

Utilities and Supporting Functions:
1. CheckDirect: Determines if the path consists of a direct connection to the end station, which is critical for handling direct connection reservations in the simulation.
2.  Djikstra: Finds the shortest path from the start to the end station using Dijkstra's algorithm, considering excluded edges and additional costs from other train paths.
3. EdgeCost: Adds a penalty to the edge cost if it’s part of another train's path, thus influencing the pathfinding algorithm.
4. Paths: Finds the fastest paths for multiple trains from start to end, considering previously excluded edges.
5. Simulate: Simulates the movement of multiple trains along their paths and prints their current positions.
6. SinglePath: Finds a single shortest path from start to end station considering excluded edges and other train paths using Dijkstra's algorithm.
7. Struck: Structures Station, Train, Connection, Edge, QueuePriority to model and solve the pathfinding problem.
8. ThereIsPath: Uses Breadth-First Search (BFS) to check if there is a connection between the start and end stations.
9. Weight: Calculates the Euclidean distance between two stations based on their coordinates. AddEdge method uses weights to create connections between stations.
