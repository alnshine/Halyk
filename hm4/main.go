package main

import (
	"fmt"
)

func main() {
	a := NewAirport()

	planes := a.Start()
	a.Close(15)

	for _, plane := range planes {
		fmt.Printf("%#v\n", plane)
	}
}
