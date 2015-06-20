// Package metro
package metro

import "io/ioutil"
import "fmt"
import "encoding/json"
import "github.com/pravj/metro/station"

func Construct(file string) {
	data, err := ioutil.ReadFile(file)

	if err == nil {
    var stations []station.Station

    err = json.Unmarshal(data, &stations)
    if err == nil {
      fmt.Println(stations)
    } else {
      panic(err)
    }

	} else {
		panic(err)
	}
}
