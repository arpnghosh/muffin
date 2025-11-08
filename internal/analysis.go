package internal

import (
	"strings"
	"unicode"

	snowballeng "github.com/kljensen/snowball/english"
)

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
