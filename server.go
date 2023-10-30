package main

import (
	"log"
	"net/http"
)

// This is heavily inspired by golang ServeMux

type MuxMethodEntry map[string]http.HandlerFunc
type MuxPathEntry map[string]MuxMethodEntry

type SmolMux struct {
	m MuxPathEntry
}

func (h *SmolMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: middleware
	log.Printf("%s %s\n", r.Method, r.URL)

	m := h.m[r.URL.Path]
	if m == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("path not registered"))
		return
	}

	f := m[r.Method]
	if f == nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not supported"))
		return
	}

	f(w, r)
}

func (h *SmolMux) GET(path string, handler http.HandlerFunc) {
	h.Handle(http.MethodGet, path, handler)
}

func (h *SmolMux) Handle(method string, path string, handler http.HandlerFunc) {
	if h.m[path] == nil {
		h.m[path] = make(MuxMethodEntry)
	}
	h.m[path][method] = handler
}
