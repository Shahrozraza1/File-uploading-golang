package routes

import (
	"ExelProject/controller"
	"github.com/gorilla/mux"
	"net/http"
)

type Routes struct{}

func New() *Routes {
	return &Routes{}
}

func (route *Routes) Init() {
	r := mux.NewRouter()
	{
		r.HandleFunc("/search", controller.SearchController).Methods("POST")
	}

	http.ListenAndServe("localhost:8088", r)
}
