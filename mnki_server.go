package main

import (
	"fmt"
	"net/http"
)

type MnkiHandlerFunc = func(req *http.Request, res http.ResponseWriter)
type MnkiEndpointMap = map[string]MnkiHandlerFunc
type MnkiMethodMap = map[string]MnkiEndpointMap

type MnkiServerMethods interface {
	Post(endpoint string, handler MnkiHandlerFunc)
	Get(endpoint string, handler MnkiHandlerFunc)
	Put(endpoint string, handler MnkiHandlerFunc)
	Patch(endpoint string, handler MnkiHandlerFunc)
	Delete(endpoint string, handler MnkiHandlerFunc)
}

type MnkiServer struct {
	handlers MnkiMethodMap
}

func (m *MnkiServer) Post(endpoint string, handler MnkiHandlerFunc) {
	m.handlers[http.MethodPost] = MnkiEndpointMap{endpoint: handler}
}

func (m *MnkiServer) Get(endpoint string, handler MnkiHandlerFunc) {
	m.handlers[http.MethodGet] = MnkiEndpointMap{endpoint: handler}
}

func (m *MnkiServer) Put(endpoint string, handler MnkiHandlerFunc) {
	m.handlers[http.MethodPut] = MnkiEndpointMap{endpoint: handler}
}

func (m *MnkiServer) Patch(endpoint string, handler MnkiHandlerFunc) {
	m.handlers[http.MethodPatch] = MnkiEndpointMap{endpoint: handler}
}

func (m *MnkiServer) Delete(endpoint string, handler MnkiHandlerFunc) {
	m.handlers[http.MethodDelete] = MnkiEndpointMap{endpoint: handler}
}

func (m *MnkiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	endpoint := r.URL.String()
	method := r.Method

	if m.handlers[method] == nil {
		http.Error(w, fmt.Sprintf("\n The method: \"%s\" is not allowed", method), http.StatusMethodNotAllowed)

	}

	if m.handlers[method][endpoint] == nil {
		http.Error(w, fmt.Sprintf("%d Route not found", http.StatusNotFound), http.StatusNotFound)

	}

	m.handlers[method][endpoint](r, w)
}

func NewMnkiServer() *MnkiServer {

	return &MnkiServer{
		handlers: MnkiMethodMap{
			http.MethodGet:    make(MnkiEndpointMap),
			http.MethodPost:   make(MnkiEndpointMap),
			http.MethodPut:    make(MnkiEndpointMap),
			http.MethodPatch:  make(MnkiEndpointMap),
			http.MethodDelete: make(MnkiEndpointMap),
		},
	}
}

/*
	package main

import (
	"net/http"
)

type MnkiServerMethods interface {
	Post(endpoint string, handler func(req *http.Request, res http.ResponseWriter))
	Get(endpoint string, handler func(req *http.Request, res http.ResponseWriter))
	Put(endpoint string, handler func(req *http.Request, res http.ResponseWriter))
	Patch(endpoint string, handler func(req *http.Request, res http.ResponseWriter))
	Delete(endpoint string, handler func(req *http.Request, res http.ResponseWriter))
}

type MnkiServer struct {
}

func (m *MnkiServer) Get(endpoint string, handler func(req *http.Request, res http.ResponseWriter)) {

	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		handler(r, w)
	})
}

func (m *MnkiServer) Post(endpoint string, handler func(req *http.Request, res http.ResponseWriter)) {

	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		handler(r, w)
	})
}
func (m *MnkiServer) Put(endpoint string, handler func(req *http.Request, res http.ResponseWriter)) {

	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		handler(r, w)
	})
}
func (m *MnkiServer) Patch(endpoint string, handler func(req *http.Request, res http.ResponseWriter)) {

	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		handler(r, w)
	})
}

func (m *MnkiServer) Delete(endpoint string, handler func(req *http.Request, res http.ResponseWriter)) {

	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		handler(r, w)
	})
}

func (m *MnkiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func NewMnkiServer() http.Handler {

	return &MnkiServer{}
}

*/
