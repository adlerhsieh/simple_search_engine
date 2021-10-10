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

func (m *MemoryStore) Set(document *entities.Document) error {
	m.data[document.ID] = document
	return nil
}

func (m *MemoryStore) Get(id string) (*entities.Document, error) {
	document, ok := m.data[id]
	if !ok {
		return nil, nil
	}

	return document, nil
}

func (m *MemoryStore) IndexByToken(token, id string) error {
	if _, ok := m.indexes[token]; !ok {
		m.indexes[token] = []string{}
	}

	// FIXME: duplicate documentID
	m.indexes[token] = append(m.indexes[token], id)

	return nil
}

func (m *MemoryStore) ListByToken(token string) ([]string, error) {
	indexes, ok := m.indexes[token]
	if !ok {
		return []string{}, nil
	}

	return indexes, nil
}
