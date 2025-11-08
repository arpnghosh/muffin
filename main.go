package main

import (
	"fmt"

	"github.com/arpnghosh/muffin/internal"
)

func main() {
	docs := []internal.Document{
		{Text: "This is a sample document about cats."},      // Doc 1: contains "cat"
		{Text: "Another document discussing cats and dogs."}, // Doc 2: contains "cat", "dog"
		{Text: "Information about dogs and birds."},          // Doc 3: contains "dog", "bird"
		{Text: "Cats birds and fish in the wild."},           // Doc 4: contains "cat", "bird", "fish"
		{Text: "Only fish in this document."},                // Doc 5: contains "fish"
		{Text: "Dogs and cats living together."},             // Doc 6: contains "dog", "cat"
		{Text: "Birds of different species."},                // Doc 7: contains "bird"
		{Text: "Empty document with common words."},
	}

	for i := range docs {
		docs[i].ID = i + 1
	}

	idx := make(internal.Index)
	idx.Add(docs)

	// Test intersection
	fmt.Println("Intersection 'cat dog':", idx.Search(internal.Intersection, "cat dog"))
	fmt.Println("Intersection 'cat bird':", idx.Search(internal.Intersection, "cat bird"))

	// Test union
	fmt.Println("Union 'cat dog':", idx.Search(internal.Union, "cat dog"))
	fmt.Println("Union 'cat bird fish':", idx.Search(internal.Union, "cat bird fish"))
}
