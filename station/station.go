// Package station
package station

type Station struct {
	Branches struct {
		A string `json:"a"`
		B string `json:"b"`
	} `json:"branches"`
	Class string `json:"class"`
	Index int64  `json:"index"`
	Name  string `json:"name"`
	Type  string `json:"type"`
}
