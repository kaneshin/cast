package app

import (
	"net/http"

	"github.com/kaneshin/cast/cmd/internal"
)

var router = internal.Router.PathPrefix("/media").Subrouter()
var render = internal.Render

func init() {
	router.HandleFunc("/", get).Methods(http.MethodGet)
	router.HandleFunc("/", post).Methods(http.MethodPost)
	router.HandleFunc("/like", post).Methods(http.MethodPost)
	router.HandleFunc("/pass", post).Methods(http.MethodPost)
}

func get(w http.ResponseWriter, r *http.Request) {
	render.Text(w, http.StatusOK, "Hello")
}

func post(w http.ResponseWriter, r *http.Request) {
	render.Text(w, http.StatusOK, "Hello")
}

func postLike(w http.ResponseWriter, r *http.Request) {
	render.Text(w, http.StatusOK, "Hello")
}

func postPass(w http.ResponseWriter, r *http.Request) {
	render.Text(w, http.StatusOK, "Hello")
}
