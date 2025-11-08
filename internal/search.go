package internal

import "sort"

// Index represents an inverted index mapping tokens to document IDs and positions
//
// Format: token -> documentID -> []positions
type Index map[string]map[int][]int

func (idx Index) Search(boolOp func([]int, []int) []int, text string) []int {
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
