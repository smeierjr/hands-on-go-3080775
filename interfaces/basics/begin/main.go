// interfaces/basics/begin/main.go
package main

import "fmt"

// define a humanoid interface with speak and walk methods returning string
type humanoid interface {
	speak()
	walk()
}

// define a person type that implements humanoid interface
type person struct{ name string }

func (p person) speak() { fmt.Printf("%s is speaking.\n", p.name) }
func (p person) walk()  { fmt.Printf("%s is walking.\n", p.name) }
func (p person) String() string {
	return fmt.Sprintf("Hello, my name is %s", p.name)
}

type dog struct{}

func (d dog) walk() { fmt.Println("Dog is walking...") }

// implement the Stringer interface for the person type

// define a dog type that can walk but not speak

func main() {
	p := person{name: "Steven"}
	doHumanThings(p)
	// invoke with a person
	// doHumanThings(person{})

	// can we invoke with a dog?
	//doHumanThings(dog{})

	fmt.Println(p)
}

func doHumanThings(h humanoid) {
	h.speak()
	h.walk()
}
