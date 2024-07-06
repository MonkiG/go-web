package main

import (
	"fmt"
	"net/http"
)

func main() {
	m := NewMnkiServer()

	m.Get("/", func(req *http.Request, res http.ResponseWriter) {
		fmt.Printf("Usando el metodo get en la ruta: %s", req.URL)
	})

	m.Post("/", func(req *http.Request, res http.ResponseWriter) {
		fmt.Printf("Usando el metodo POST en la ruta %s", req.URL)
	})

	http.ListenAndServe(":8080", m)
}
