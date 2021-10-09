package memorystore

import "github.com/adlerhsieh/simple_search_engine/entities"

type MemoryStore struct {
	data    map[string]*entities.Document
	indexes map[string][]string
}

func New() *MemoryStore {
	return &MemoryStore{
		data:    map[string]*entities.Document{},
		indexes: map[string][]string{},
	}
}

func (m *MemoryStore) SaveDocument(document *entities.Document) error {
	// for _, token := range tokens {
	// 	if _, ok := m.data[token]; !ok {
	// 		m.indexes[token] = []string{}
	// 	}
	//
	// 	// FIXME: duplicate documentID
	// 	m.indexes[token] = append(m.indexes[token], documentID)
	// }

	return nil
}

func (m *MemoryStore) ListDocumentsByToken(token string) []string {
	return []string{}
}
