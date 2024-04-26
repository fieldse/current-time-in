// Functions to load the JSON city data from ../data directory
package citylookup

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/fieldse/current-time-in/shared"
)

var (
	DataDirectory string
	DataJSONFile  string
)

// Set the data directory
func init() {
	rootDir := getProjectRootDir()
	p := path.Join(rootDir, "data")
	shared.ExistsOrPanic(p) // validate the data directory exists or panic
	DataDirectory = p
	DataJSONFile = path.Join(p, "cityMap.json")
}

// Return path to the project root directory --
func getProjectRootDir() string {
	curDir, _ := os.Getwd()
	return path.Dir(path.Dir(curDir))
}

// readCityData reads the raw data from cityMap.json to bytes
func readCityData() ([]byte, error) {
	b, err := os.ReadFile(DataJSONFile)
	if err != nil {
		log.Printf("failed to read cityMap.json: %v", err)
	}
	return b, err
}

// loadCityData loads the city data from cityMap.json
func loadCityData() ([]CityData, error) {
	var cityData []CityData
	b, err := readCityData()
	if err != nil {
		return []CityData{}, fmt.Errorf("failed to read city data: %w", err)
	}
	err = json.Unmarshal(b, &cityData)
	if err != nil {
		return []CityData{}, fmt.Errorf("error parsing city data: %w", err)
	}
	return cityData, nil
}
