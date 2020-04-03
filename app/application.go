package app

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrls()

	server := &http.Server{
		Handler: router,
		Addr: "127.0.0.1:8000",
		//Timeouts
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout: 2 * time.Second,
		IdleTimeout: 60 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}