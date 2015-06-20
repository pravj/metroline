// Package controller supervises the construction
// For example, which station to build next and on which branch/line
package controller

import (
  "fmt"
  "math/rand"
  "github.com/pravj/metro/station"
  "github.com/pravj/metro/branch"
  "github.com/pravj/metro/git"
)

// Controller struct with its required fields
type Controller struct {
  Stations []station.Station
  Inventory map[string][]int
  Constraints map[string]bool

  CurrentBranch *branch.Branch
  Branches []string
  BranchManager map[string]*branch.Branch

  IsStarted bool
  IsEnded bool
}

// New returns a new controller instance
func New(stations []station.Station, inventory map[string][]int, constraints map[string]bool) *Controller {
  return &Controller{Stations: stations, Inventory: inventory, Constraints: constraints, BranchManager: make(map[string]*branch.Branch), IsStarted: false, IsEnded: false}
}

// Control takes over the construction here
func (c *Controller) Control() {
  c.collectBranches()
  git.GitInit()

  branchName := c.getLeaderBranch()
  c.setupBranch(branchName)

  for _, index := range c.Inventory[c.CurrentBranch.Name] {
    fmt.Println(c.Stations[index])
  }
}

// collectBranch creates a slice of all the branches(lines) in the network
func (c *Controller) collectBranches() {
  for branch := range c.Inventory {
    c.Branches = append(c.Branches, branch)
  }
}

// setupBranch sets a branch to work on
func (c *Controller) setupBranch(branchName string) {
  newBranch := branch.New(branchName, c.Inventory[branchName], &c.Stations)

  c.BranchManager[branchName] = newBranch
  c.CurrentBranch = newBranch

  newBranch.Start()
}

// getLeaderBranch returns a branch name which can be started as an orphan(independent)
func (c *Controller) getLeaderBranch() string {
  index := rand.Intn(len(c.Branches))
  branchName := c.Branches[index]

  if c.Constraints[branchName] {
    return branchName
  } else {
    return c.getLeaderBranch()
  }
}

// branch manager : name to branch mapping
