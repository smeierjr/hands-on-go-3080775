// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	title string
	author
}

// define a library type to track a list of books
type library map[string][]book

// define addBook to add a book to the library
func (lib library) addBook(b book) {
	lib[b.name] = append(lib[b.name], b)
}

// define a lookupByAuthor function to find books by an author's name
//
func (lib library) lookupByAuthor(name string) []book {
	return lib[name]
}

func main() {

	// create a new library
	lib := library{}

	// add 2 books to the library by the same author

	author1 := author{name: "Stephen King"}
	lib.addBook(book{
		title:  "The Shining",
		author: author1,
	})

	lib.addBook(book{
		title:  "Carrie",
		author: author1,
	})

	// add 1 book to the library by a different author
	//

	lib.addBook(book{
		title: "Rainbow Six",
		author: author{
			name: "Tom Clancy",
		},
	})

	// dump the library with spew
	spew.Dump(lib)

	// lookup books by known author in the library
	//
	books := lib.lookupByAuthor("Stephen King")
	spew.Dump(books)

	// print out the first book's title and its author's name
	//
	if len(books) != 0 {
		fmt.Printf("%v by %v\n", books[0].title, books[0].author.name)
	}
}
