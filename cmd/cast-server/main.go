package main

import (
	"log"
	"net/http"

	_ "github.com/kaneshin/cast/cmd/cast-server/media"
	"github.com/kaneshin/cast/cmd/internal"
)

func main() {
	http.Handle("/", internal.Router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
