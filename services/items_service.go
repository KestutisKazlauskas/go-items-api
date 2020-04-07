package services

import (
	"github.com/KestutisKazlauskas/go-utils/rest_errors"
	"github.com/KestutisKazlauskas/go-items-api/domain/items"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, *rest_errors.RestErr)
	Get(string) (*items.Item, *rest_errors.RestErr)
}

type itemsService struct {}

func (s *itemsService) Create(itemRequest items.Item) (*items.Item, *rest_errors.RestErr) {
	if err := itemRequest.Save(); err != nil {
		return nil, err
	}

	return &itemRequest, nil
}

func(s *itemsService) Get(id string) (*items.Item, *rest_errors.RestErr) {
	return nil, rest_errors.NewNotFoundError("Not Found")
}