package router

import "github.com/gorilla/mux"

var router *mux.Router

func New() *mux.Router {
	router = mux.NewRouter()
	mapUrls()
	return router
}
