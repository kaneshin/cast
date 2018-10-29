package internal

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var (
	// Router is.
	Router *mux.Router

	// Render is.
	Render *render.Render

	// EchoHandler write echo message into response.
	EchoHandler = func(msg string) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			Render.Text(w, http.StatusOK, msg)
		}
	}
)

func init() {
	Router = mux.NewRouter()
	Router = Router.StrictSlash(true)

	opt := render.Options{}
	Render = render.New(opt)
}

func init() {
	Router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		Render.Text(w, http.StatusOK, "pong")
	}).Methods(http.MethodGet)
}
