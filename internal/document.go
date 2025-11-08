package internal

type Document struct {
	Text string
	ID   int
}

func (idx Index) Add(docs []Document) {
	for _, doc := range docs {
		for position, token := range analyze(doc.Text) {
			if idx[token] == nil {
				idx[token] = make(map[int][]int)
			}
			idx[token][doc.ID] = append(idx[token][doc.ID], position)
		}
	}
}
