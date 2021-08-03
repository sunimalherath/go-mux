package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sunimalherath/go-mux/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/products", handlers.ProductHandler)
	r.HandleFunc("/articles", handlers.ArticlesHandler)
	http.Handle("/", r)
}
