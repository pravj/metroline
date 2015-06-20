// Package parser implements parsing of git command outputs.
// Then it collects vital information from them, like SHA-1 hash of recent commit etc.
package parser

import (
  "fmt"
  "regexp"
)

var (
  HashPattern string
)

func init() {
  HashPattern = " [a-z0-9]{7}"
}

// CommitHash returns the SHA-1 commit hash string for a commit
func CommitHash(output string) (err error, hash string) {
  match, err := regexp.MatchString(HashPattern, output)

  if match {
    r, _ := regexp.Compile(HashPattern)

    return nil, r.FindString(output)
  } else {
    return err, nil
  }
}
