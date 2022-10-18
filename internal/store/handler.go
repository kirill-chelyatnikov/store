package store

import (
	"fmt"
	"net/http"
)

type handler struct {
}

func (h handler) Register() {
	http.HandleFunc("/items", getAll)
}

func getAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello!")
}

func NewHandler() *handler {
	return &handler{}
}
