package internal

// PhraseSearch searches for an exact phrase in the index.
// It returns the documents where all tokens appear consecutively.
//
// Example:
//
// phrase: "cats and dog"
//
// tokens: ["cat", "dog"]
//
// Must find: documents where "cat" is at position N and "dog" is at pos N+1
func (idx Index) PhraseSearch(text string) []int {
	result := make(map[string]map[int][]int)

	docIds := idx.Search(Intersection, text)
	tokens := analyze(text)

	for _, token := range tokens {
		for _, docId := range docIds {
			if position_array, ok := idx[token][docId]; ok {
				if result[token] == nil {
					result[token] = make(map[int][]int)
				}
				result[token][docId] = append(result[token][docId], position_array...)
			}
		}
	}
	var matchingDocs []int

	for _, docId := range docIds {
		firstTokenPositions := result[tokens[0]][docId]

		for _, startPos := range firstTokenPositions {
			matched := true

			for i := 1; i < len(tokens); i++ {
				expectedPos := startPos + i
				tokenPositions := result[tokens[i]][docId]
				found := false
				for _, pos := range tokenPositions {
					if pos == expectedPos {
						found = true
						break
					}
					if !found {
						matched = false
						break
					}
				}
			}
			if matched {
				matchingDocs = append(matchingDocs, docId)
				break
			}
		}
	}
	return matchingDocs
}
