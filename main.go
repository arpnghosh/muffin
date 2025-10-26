package main

import (
	"fmt"
	"strings"
	"unicode"

	snowballeng "github.com/kljensen/snowball/english"
)

type document struct {
	Text string
	ID   int
}

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

var stopwords = map[string]struct{}{
	"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
	"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
}

func stopwordFilter(tokens []string) []string {
	r := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if _, ok := stopwords[token]; !ok {
			r = append(r, token)
		}
	}
	return r
}

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
	tokens = stopwordFilter(tokens)
	tokens = stemmerFilter(tokens)
	return tokens
}

type index map[string][]int

func (idx index) add(docs []document) {
	for _, doc := range docs {
		for _, token := range analyze(doc.Text) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				continue
			}
			idx[token] = append(ids, doc.ID)
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

func (idx index) search(text string) []int {
	var r []int
	for _, token := range analyze(text) {
		if ids, ok := idx[token]; ok {
			if r == nil {
				r = ids
			} else {
				r = intersection(r, ids)
			}
		} else {
			return nil
		}
	}
	return r
}

func main() {
	docs := []document{
		{Text: "This is a sample document about cats."},
		{Text: "Another document discussing cats pets."},
		{Text: "Information on birds discussing and their."},
	}

	for i := range docs {
		docs[i].ID = i + 1
	}

	idx := make(index)
	idx.add(docs)

	results := idx.search("cat")
	if len(results) == 0 {
		fmt.Println("search not found")
	} else {
		fmt.Println("results found in document:", results)
	}
}
