package main

import (
	"github.com/adlerhsieh/simple_search_engine/entities"
	"github.com/adlerhsieh/simple_search_engine/indexer"
	"github.com/adlerhsieh/simple_search_engine/memorystore"
)

func main() {
	documents := []*entities.Document{
		{
			ID:      "1",
			Content: "I have a book",
		},
		{
			ID:      "2",
			Content: "This is a book",
		},
		{
			ID:      "3",
			Content: "Back to the inverted index",
		},
	}

	store := memorystore.New()
	appIndexer := indexer.New(store)

	for _, document := range documents {
		appIndexer.Index(document)
	}
}
