// types/maps/initialization/begin/main.go
package main

import "fmt"

type author struct {
	name string
}

func main() {
	// declare a map of string keys and author values

	// initialize the map with make
	authors := map[string]author{
		"tm": {name: "Toni Morrison"},
		"ma": {name: "Marcus Aurelius"},
	}

	// add authors to the map

	// print the map with fmt.Printf
	// fmt.Printf("%#v\n", authors)

	// read a value from the map with a known key
	fmt.Printf("%#v\n", authors["tm"])
}
