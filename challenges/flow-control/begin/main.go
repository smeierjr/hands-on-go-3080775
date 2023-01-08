// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recoved from panic: ", err)
		}

	}()

	// use os package to read the file specified as a command line argument
	//

	//data, _ := os.ReadFile("/workspaces/hands-on-go-3080775/challenges/data.txt")
	bs, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("failed to read file: %s", os.Args[1]))
	}
	//fmt.Println(string(fileBytes))

	// convert the bytes to a string
	//
	data := string(bs)

	// initialize a map to store the counts
	analysis := make(map[string]int)

	// use the standard library utility package that can help us split the string into words
	words := strings.Fields(data)

	// capture the length of the words slice
	analysis["words"] = len(words)
	spew.Dump(analysis)
	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	for _, word := range words {
		for _, char := range word {
			if unicode.IsLetter(char) {
				analysis["letters"]++
			} else if unicode.IsNumber(char) {
				analysis["numbers"]++
			} else {
				analysis["symbols"]++
			}
		}
	}
	// dump the map to the console using the spew package
	spew.Dump(analysis)

}
