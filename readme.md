# STATIONS PATHFINDER

A Go command-line application that finds efficient train routes between two stations and simulates train movement turn by turn.

## Overview

This project solves a railway pathfinding problem where multiple trains must travel from a start station to an end station using the minimum number of movement turns.

The program reads a network map file containing stations and connections, validates the input, finds suitable paths, and prints train movements for each turn while respecting constraints like fixed-block signaling and station occupancy.

## Features

- Parses railway network maps from file
- Validates stations, connections, and input arguments
- Finds paths using Dijkstra’s algorithm and BFS
- Supports multiple trains moving simultaneously
- Prevents collisions using track and station occupancy rules
- Simulates train movement turn by turn
- Handles invalid input gracefully

## Technologies Used

- Go
- Graph data structures
- Dijkstra’s algorithm
- Breadth-First Search (BFS)
- Command-line interface (CLI)

## Usage

go run . [path to file containing network map] [start station] [end station] [number of trains]

Example:
go run . network.map waterloo st_pancras 4

## Network Map Format

Stations:
stations:
waterloo,3,1
victoria,6,7
euston,11,23
st_pancras,5,15

Connections:
connections:
waterloo-victoria
waterloo-euston
st_pancras-euston
victoria-st_pancras

Comments (#), blank lines, and extra whitespace are ignored.

## Example Output

T1-victoria T2-euston
T1-st_pancras T2-st_pancras T3-victoria T4-euston
T3-st_pancras T4-st_pancras

Each line represents one movement turn.

## Project Flow

1. Load and parse the network map
2. Validate input and structure
3. Check if a valid path exists
4. Calculate paths for trains
5. Simulate movement step by step
6. Output results

## Key Components

- Pathfinding using Dijkstra’s algorithm
- BFS for connectivity checks
- Simulation engine for train movement
- Edge cost adjustments to avoid conflicts

## Error Handling

Handles cases such as:
- invalid arguments
- missing or duplicate stations
- invalid connections
- no path between stations
- invalid coordinates

Outputs Error and exits safely.

## Testing

Run:
make

This generates testresults.txt with predefined test outputs.

## Learning Outcomes

- Graph-based pathfinding implementation
- Multi-agent movement simulation
- CLI application development in Go
- Handling large datasets efficiently
