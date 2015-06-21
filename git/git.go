// Package git implements a tiny wrapper around git in golang.
// It lets you create branches, commit into them and merge them.
//
// It uses golang's "os/exec" module for executing commands.
// Then it parse the command output for further use.
package git

import (
	"fmt"
	"github.com/pravj/metroline/git/parser"
	"os/exec"
)

var (
	cmdOutput []byte
	cmdErr    error
	mainCmd   string
)

func init() {
	mainCmd = "git"
}

// GitInit initiates an empty git repository in the current directory
func GitInit() (err error) {
	args := []string{"init"}

	cmdErr := exec.Command(mainCmd, args...).Run()
	if cmdErr == nil {
		// this commit will go into the master branch
		GitCommit("Initial Commit")
		fmt.Println("Initial Commit")
	}

	return cmdErr
}

// GitRenameBranch renames a given branch
func GitRenameBranch(oldName, newName string) (err error) {
	args := []string{"branch", "-m", oldName, newName}

	return exec.Command(mainCmd, args...).Run()
}

// GitCommit does an empty(no content) commit with a specified commit message and returns its SHA-1 hash
func GitCommit(message string) (err error, hash string) {
	args := []string{"commit", "--allow-empty", "-m", message}
	cmdOutput, cmdErr = exec.Command(mainCmd, args...).Output()

	if cmdErr != nil {
		return cmdErr, ""
	} else {
		return parser.CommitHash(string(cmdOutput))
	}
}

// GitCreateBranch creates a new branch, can also create an 'orphan' branch if specified
func GitCreateBranch(name string, isOrphan bool) (err error) {
	args := []string{"checkout", "-b", name}
	if isOrphan {
		args[1] = "--orphan"
	}

	return exec.Command(mainCmd, args...).Run()
}

// GitMerge merges a given branch into current branch, with a specified commit message
func GitMerge(branchName, message string) (err error, output string) {
	args := []string{"merge", branchName, "-m", message}

	cmdOutput, cmdErr = exec.Command(mainCmd, args...).Output()
	return cmdErr, string(cmdOutput)
}

// GitCheckout switches to a specified branch from current branch
func GitCheckout(name string) (err error, output string) {
	args := []string{"checkout", name}

	cmdOutput, cmdErr = exec.Command(mainCmd, args...).Output()
	return cmdErr, string(cmdOutput)
}

// GitResetBranch forcefully checkout a given branch to current HEAD
func GitResetBranch(branch, commit string) (err error) {
	args := []string{"branch", "-f", branch, "HEAD"}

	return exec.Command(mainCmd, args...).Run()
}
