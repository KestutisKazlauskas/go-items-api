package items

import (
	"github.com/KestutisKazlauskas/go-utils/rest_errors"
	"github.com/KestutisKazlauskas/go-items-api/clients/elasticsearch"
)

const (
	indexItems = "items"
)

// Database only here
func (i *Item) Save() *rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("Error already loged", nil, nil)
	}
	// could generate myself from db for example
	i.Id = result.Id
	return nil
}