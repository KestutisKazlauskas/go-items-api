package controllers

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/KestutisKazlauskas/go-utils/rest_errors"
	"github.com/KestutisKazlauskas/go-oauth/oauth"
	"github.com/KestutisKazlauskas/go-items-api/domain/items"
	"github.com/KestutisKazlauskas/go-items-api/services"
	"github.com/KestutisKazlauskas/go-items-api/utils/http_utils"
	"github.com/KestutisKazlauskas/go-utils/logger"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemsController struct {}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.Authenticate(r); err != nil {
		http_utils.RespondError(w, err)
		return 
	}

	logger.Log.Info(r.Header.Get("X-User-Id"))

	sellerId := oauth.GetUserId(r)
	if sellerId == 0 {
		//Error or not authenticated reqeust
		respErr := rest_errors.NewBadRequestError("invalid access token")
		http_utils.RespondError(w, respErr)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if marshalErr := json.Unmarshal(requestBody, &itemRequest); marshalErr != nil {
		respErr := rest_errors.NewBadRequestError("invalid item json body")
		http_utils.RespondError(w, respErr)
		return
	}

	itemRequest.Seller = sellerId

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}

	http_utils.RespondJson(w, http.StatusOK, result)

}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}