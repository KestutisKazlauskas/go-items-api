package app

import (
	"net/http"
	"github.com/KestutisKazlauskas/go-items-api/controllers"
)

func mapUrls() {
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
}