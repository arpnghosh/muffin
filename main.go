package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	snowballeng "github.com/kljensen/snowball/english"
)

type document struct {
	Text string
	ID   int
}

// tokenize splits text into searchable tokens
// by breaking on non-alphanumeric characters.
// It preserves both letters and numbers, making
// it suitable for full-text search indexing.
//
// Example:
//
//	input:  "Hello, World! User123"
//	output: ["Hello", "World", "User123"]
func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsNumber(r) && !unicode.IsLetter(r)
	})
}

func lowercaseFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}

/*
* The first commit contains the stopwords filter.
* After updating my inverted index type
* to store the position of each token,
* I had to remove the stopword filter
 */

// var stopwords = map[string]struct{}{
// 	"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
// 	"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
// }

// func stopwordFilter(tokens []string) []string {
// 	r := make([]string, 0, len(tokens))
// 	for _, token := range tokens {
// 		if _, ok := stopwords[token]; !ok {
// 			r = append(r, token)
// 		}
// 	}
// 	return r
// }

func stemmerFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = snowballeng.Stem(token, false)
	}
	return r
}

func analyze(text string) []string {
	tokens := tokenize(text)
	tokens = lowercaseFilter(tokens)
	// tokens = stopwordFilter(tokens)
	tokens = stemmerFilter(tokens)
	return tokens
}

type index map[string]map[int][]int

func (idx index) add(docs []document) {
	for _, doc := range docs {
		for position, token := range analyze(doc.Text) {
			if idx[token] == nil {
				idx[token] = make(map[int][]int)
			}
			idx[token][doc.ID] = append(idx[token][doc.ID], position)
		}
	}
}

func intersection(a []int, b []int) []int {
	minLen := min(len(a), len(b))
	r := make([]int, 0, minLen)

	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

func union(a []int, b []int) []int {
	r := make([]int, 0, len(a)+len(b))

	var i, j int

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			r = append(r, a[i])
			i++
		} else if a[i] > b[j] {
			r = append(r, b[j])
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}

	for i < len(a) {
		r = append(r, a[i])
		i++
	}

	for j < len(b) {
		r = append(r, b[j])
		j++
	}

	return r
}

func (idx index) search(boolOp func([]int, []int) []int, text string) []int {
	var r []int
	for _, token := range analyze(text) {
		if innerMap, ok := idx[token]; ok {
			ids := make([]int, 0, len(innerMap))
			for docId := range innerMap {
				ids = append(ids, docId)
			}
			sort.Ints(ids)
			if r == nil {
				r = ids
			} else {
				r = boolOp(r, ids)
			}
		}
	}
	return r
}

func main() {
	docs := []document{
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

	idx := make(index)
	idx.add(docs)

	// Test intersection
	fmt.Println("Intersection 'cat dog':", idx.search(intersection, "cat dog"))
	fmt.Println("Intersection 'cat bird':", idx.search(intersection, "cat bird"))

	// Test union
	fmt.Println("Union 'cat dog':", idx.search(union, "cat dog"))
	fmt.Println("Union 'cat bird fish':", idx.search(union, "cat bird fish"))
}
