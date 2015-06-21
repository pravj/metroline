// Package branch
package branch

import "fmt"
import "github.com/pravj/metroline/station"
import "github.com/pravj/metroline/git"

type Branch struct {
	Name string

	Stations []station.Station
	StationsIndex []int
	Current  int

	HEAD string

	IsWaiting   bool
	IsComplete bool

	Messenger chan string
}

func New(name string, stationsIndex []int, stations *[]station.Station, messenger chan string) *Branch {
	return &Branch{Name: name, Stations: *stations, StationsIndex: stationsIndex, Messenger: messenger}
}

func (b *Branch) Start() {
	b.Messenger <- fmt.Sprintf("starting branch %v", b.Name)

	err := git.GitCreateBranch(b.Name, true)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("created branch")
	}

	b.BuildStation()
}

func (b *Branch) BuildStation() {
	newStation := b.Stations[b.StationsIndex[b.Current]]

	b.Messenger <- fmt.Sprintf("building station %v", newStation.Name)

	switch newStation.Type {
	case "simple":
		git.GitCommit(newStation.Name)
	case "two":
		b.Messenger <- "two"
		return
	case "twi":
		b.Messenger <- "twi"
		return
	case "fwc":
		b.Messenger <- "fwc"
		return
	}

	b.Current += 1
	b.BuildStation()
}
