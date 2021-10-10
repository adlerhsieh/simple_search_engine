package interfaces

import "github.com/adlerhsieh/simple_search_engine/entities"

type DataStorage interface {
	Set(document *entities.Document) error
	Get(id string) (*entities.Document, error)
	IndexByToken(token, id string) error
	ListByToken(token string) ([]string, error)
}
