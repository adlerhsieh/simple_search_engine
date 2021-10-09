package indexer

import (
	"github.com/adlerhsieh/simple_search_engine/interfaces"
)

type Indexer struct {
	Storage interfaces.DataStorage
}

func New(storage interfaces.DataStorage) *Indexer {
	return &Indexer{
		Storage: storage,
	}
}

func (i *Indexer) Index() error {
	// err := i.Storage.SaveDocument(tokens, documentID)
	// if err != nil {
	// 	return err
	// }
	return nil
}
