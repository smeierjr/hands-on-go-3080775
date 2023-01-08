// types/structs/fields/begin/main.go
package main

import "fmt"

// define a struct type for author
//
type author struct {
	first, last string
}

func (a author) fullName() string {
	return a.first + " " + a.last
}

func main() {
	// intialize author
	a := author{
		first: "Clark",
		last:  "Kent",
	}

	// print the author
	fmt.Printf("%#v\n", a)
	fmt.Println(a.fullName())
}
