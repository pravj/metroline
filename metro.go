// Package metro
package metro

import "io/ioutil"
import "encoding/json"
import "github.com/pravj/metro/station"
import "github.com/pravj/metro/controller"

// stations, branch storage as a struct to unmarshal json
type stationStore struct {
  Name string `json:"name"`
  Stations []int `json:"stations"`
	Leader bool `json:"leader"`
}

// mapStationData returns a map from string(station) to int(index of stations)
func mapStationData(mapFile string) (map[string][]int, map[string]bool) {
  mapData, err := ioutil.ReadFile(mapFile)

	var storage []stationStore
	inventory := make(map[string][]int)
	constraints := make(map[string]bool)

	if err == nil {
		errNew := json.Unmarshal(mapData, &storage)
		if errNew == nil {

		  for _, entity := range storage {
				constraints[entity.Name] = entity.Leader
		    for _, entityStation := range entity.Stations {
		      inventory[entity.Name] = append(inventory[entity.Name], entityStation)
		    }
		  }

			return inventory, constraints
		} else {
			panic(errNew)
		}
	} else {
		panic(err)
	}
}

// Construct initiates the metro line construction
// It reads the json data for stations and hand it over to higher authorities/departments
func Construct(stationFile, mapFile string) {
	stationData, err := ioutil.ReadFile(stationFile)
	if err == nil {
    var stations []station.Station

    errNew := json.Unmarshal(stationData, &stations)
    if errNew == nil {
			inventory, constraints := mapStationData(mapFile)

			// controller takes over the construction here
	    controller.New(stations, inventory, constraints).Control()
    } else {
      panic(errNew)
    }

	} else {
		panic(err)
	}
}
