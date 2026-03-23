package reader

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"stations/pkg/pathfinding"
	"strconv"
	"strings"
)

// Used to track duplicates
var nameMap = make(map[string]bool)
var coordinatesMap = make(map[string]bool)

// ParseInputFile reads the map file and populates global slices with station and connection data
func ParseInputFile(FilePath string) (map[string]*pathfinding.Station, map[string][]string, []error) {
	file, err := os.Open(FilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file: %v\n", err)
		return nil, nil, []error{err}
	}
	defer file.Close()

	var parseErrs []error

	// First check that the necessary sections are present
	if err := CheckStationsConnections(file); err != nil {
		fmt.Fprintf(os.Stderr, "problem parsing file: %v\n", err)
		parseErrs = append(parseErrs, err)
		return nil, nil, parseErrs
	}

	scanner := bufio.NewScanner(file)

	ParsingStations := false
	ParsingConnections := false
	NumberOfStations := 0

	stations := make(map[string]*pathfinding.Station)
	connections := make(map[string][]string)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip blank lines and lines that start with '#' (considering leading whitespace)
		if line == "" || strings.HasPrefix(strings.TrimSpace(line), "#") {
			continue // Skip blank lines and comments
		}

		if strings.HasPrefix(line, "stations:") {
			ParsingStations = true
			ParsingConnections = false
			continue
		}

		if strings.HasPrefix(line, "connections:") {
			ParsingStations = false
			ParsingConnections = true
			continue
		}

		if ParsingStations {
			err = CheckStationLine(line)
			if err != nil {
				parseErrs = append(parseErrs, err)
			} else {
				name, x, y, err := ParseStationLine(line)
				if err != nil {
					parseErrs = append(parseErrs, err)
				} else {
					// Create a new Station struct and add it to the stations map
					stations[name] = &pathfinding.Station{
						Name: name,
						X:    x,
						Y:    y,
					}
				}
			}
			NumberOfStations++
			if NumberOfStations > 10000 {
				err := fmt.Errorf("%s contains more than 10,000 stations", FilePath)
				parseErrs = append(parseErrs, err)
			}
		}

		if ParsingConnections {
			err := ParseConnection(line, stations, connections)
			if err != nil {
				parseErrs = append(parseErrs, err)
				continue
			}
		}
	}

	if err := scanner.Err(); err != nil {
		parseErrs = append(parseErrs, err)
	}

	return stations, connections, parseErrs
}

func CheckStationsConnections(file *os.File) error {
	scanner := bufio.NewScanner(file)

	FileHasStations := false
	FileHasConnections := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip blank lines and lines that start with '#'
		if line == "" || strings.HasPrefix(strings.TrimSpace(line), "#") {
			continue // Skip blank lines and comments
		}
		if strings.HasPrefix(line, "stations:") {
			FileHasStations = true
		}
		if strings.HasPrefix(line, "connections:") {
			FileHasConnections = true
		}
		// Exit if both sections are found
		if FileHasStations && FileHasConnections {
			break
		}
	}

	// Reset the file pointer to the beginning (learned this in the first task!)
	_, err := file.Seek(0, 0)
	if err != nil {
		err = errors.New("error resetting file pointer")
		return err
	}

	// Check if both sections were found
	if !FileHasStations && !FileHasConnections {
		err := errors.New("stations and connections not found in file")
		return err
	} else if !FileHasStations {
		err := errors.New("stations section not found in file")
		return err
	} else if !FileHasConnections {
		err := errors.New("connections section not found in file")
		return err
	}

	return nil
}

// CheckStationLine validates the format of a station line
func CheckStationLine(line string) error {
	// Updated regex to allow optional spaces around commas and handle comments
	StationRegex := regexp.MustCompile(`^[a-z0-9_]+\s*,\s*[0-9]+\s*,\s*[0-9]+(?:\s*#.*)?$`)
	if !StationRegex.MatchString(line) {
		err := fmt.Errorf("invalid station format: %s", line)
		return err
	}
	return nil
}

// CheckConnectionLine validates the format of a connection line
func CheckConnectionLine(line string) error {
	ConnectionRegex := regexp.MustCompile(`^[a-z0-9_]+-[a-z0-9_]+$`)
	if !ConnectionRegex.MatchString(line) {
		err := fmt.Errorf("invalid connection format: %s", line)
		return err
	}
	return nil
}

// ParseStationLine extracts station details from a line
func ParseStationLine(line string) (string, int, int, error) {
	// Split the line on '#' to remove comments
	line = strings.SplitN(line, "#", 2)[0]

	// Split the station info by commas
	StationData := strings.Split(line, ",")
	if len(StationData) != 3 {
		return "", 0, 0, fmt.Errorf("invalid station format: '%s'", line)
	}

	// Extract and trim spaces around each part
	ParsedName := strings.TrimSpace(StationData[0])
	ParsedXStr := strings.TrimSpace(StationData[1])
	ParsedYStr := strings.TrimSpace(StationData[2])

	// Convert trimmed strings to integers
	ParsedX, err := strconv.Atoi(ParsedXStr)
	if err != nil {
		return "", 0, 0, fmt.Errorf("invalid X coordinate '%s' in entry: '%s'", ParsedXStr, line)
	}
	ParsedY, err := strconv.Atoi(ParsedYStr)
	if err != nil {
		return "", 0, 0, fmt.Errorf("invalid Y coordinate '%s' in entry: '%s'", ParsedYStr, line)
	}

	// Check for duplicate station names
	if _, exists := nameMap[ParsedName]; exists {
		return "", 0, 0, fmt.Errorf("duplicate station name '%s' found", ParsedName)
	}
	nameMap[ParsedName] = true

	// Check for duplicate coordinates
	coordKey := fmt.Sprintf("%d,%d", ParsedX, ParsedY)
	if _, exists := coordinatesMap[coordKey]; exists {
		if !coordinatesMap[coordKey] {
			err := fmt.Errorf("duplicate coordinates (%d, %d) found", ParsedX, ParsedY)
			coordinatesMap[coordKey] = true // Mark this coordinate as reported
			return "", 0, 0, err
		}
	} else {
		coordinatesMap[coordKey] = false // False means not reported yet
	}

	return ParsedName, ParsedX, ParsedY, nil
}

func ParseConnection(line string, stations map[string]*pathfinding.Station, connections map[string][]string) error {
	// Split the connection info by dash
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return fmt.Errorf("invalid connection format: %s", line)
	}

	station1 := strings.TrimSpace(parts[0])
	station2 := strings.TrimSpace(parts[1])

	// Check if both stations exist
	if _, ok := stations[station1]; !ok {
		return fmt.Errorf("connection to invalid station(s): %s", line)
	}
	if _, ok := stations[station2]; !ok {
		return fmt.Errorf("connection to invalid station(s): %s", line)
	}

	if HasConnection(connections[station1], station2) {
		return fmt.Errorf("duplicate connection: %s-%s", station1, station2)
	}

	connections[station1] = append(connections[station1], station2)
	connections[station2] = append(connections[station2], station1)

	return nil
}

// Helper function to check if a connection already exists
func HasConnection(connections []string, station string) bool {
	for _, conn := range connections {
		if conn == station {
			return true
		}
	}
	return false
}
