// concurrency/sync/begin/main.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	// given a list of names
	names := []string{"Steven", "Priscilla", "Maddox"}

	// initialize a map to store the length of each name
	namesMap := make(map[string]int)

	// initialize a wait group for the goroutines that will be launched
	var wg sync.WaitGroup
	wg.Add(len(names))
	var mu sync.Mutex

	// launch a goroutine for each name we want to process
	for _, name := range names {
		go func(name string) {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			namesMap[name] = len(name)
		}(name)
	}

	// wait for all goroutines to finish
	wg.Wait()

	// print the map
	fmt.Println(namesMap)
}
