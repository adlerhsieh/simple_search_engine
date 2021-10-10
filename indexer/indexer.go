package indexer

import (
	"github.com/adlerhsieh/simple_search_engine/entities"
	"github.com/adlerhsieh/simple_search_engine/interfaces"
	"github.com/adlerhsieh/simple_search_engine/tokenizer"
)

type Indexer struct {
	Storage interfaces.DataStorage
}

func New(storage interfaces.DataStorage) *Indexer {
	return &Indexer{
		Storage: storage,
	}
}

func (i *Indexer) Index(document *entities.Document) error {
	err := i.Storage.Set(document)
	if err != nil {
		return err
	}

	tokens, err := tokenizer.Tokenize(document.Content)
	if err != nil {
		return err
	}

	for _, t := range tokens {
		err := i.Storage.IndexByToken(t, document.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
