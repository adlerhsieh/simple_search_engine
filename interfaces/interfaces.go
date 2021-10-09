package interfaces

import "github.com/adlerhsieh/simple_search_engine/entities"

type DataStorage interface {
	SaveDocument(document *entities.Document) error
	ListDocumentsByToken(token string) []string
}
