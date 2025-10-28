# Muffin

<p align="center">
  <img src="public/logo.png" alt="Muffin Logo" width="200">
</p>

## Project Description

Muffin is a simple full-text search engine implemented in Go. It builds an inverted index from documents, supporting tokenization, stemming, and boolean search operations (intersection for AND, union for OR).

## Features

- **Tokenization**: Splits text into searchable tokens, preserving letters and numbers.
- **Stemming**: Uses the Snowball English stemmer to reduce words to their root forms.
- **Inverted Index**: Stores term positions within documents for efficient searching.
- **Boolean Search**: Supports intersection (AND) and union (OR) operations on search queries.

## Usage

### Adding Documents and Building Index

```go
docs := []document{
    {Text: "This is a sample document about cats."},
    {Text: "Another document discussing cats and dogs."},
    // ... more documents
}

for i := range docs {
    docs[i].ID = i + 1
}

idx := make(index)
idx.add(docs)
```

### Searching

```go
// Intersection (AND) search
results := idx.search(intersection, "cat dog")
fmt.Println("Documents containing both 'cat' and 'dog':", results)

// Union (OR) search
results = idx.search(union, "cat bird fish")
fmt.Println("Documents containing 'cat', 'bird', or 'fish':", results)
```

Run the program with `go run main.go` to see example searches on sample documents.
