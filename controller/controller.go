// Package controller supervises the construction
// For example, which station to build next and on which branch/line
package controller

import (
  "fmt"
  "github.com/pravj/metro/station"
  //"github.com/pravj/metro/branch"
)

// Controller struct with its required fields
type Controller struct {
  Stations []station.Station
  Inventory map[string][]int
}

// New returns a new controller instance
func New(stations []station.Station, inventory map[string][]int) *Controller {
  return &Controller{Stations: stations, Inventory: inventory}
}

// Control takes over the construction here
func (c *Controller) Control() {
  fmt.Println(c.Inventory)
}

// branch manager : name to branch mapping
