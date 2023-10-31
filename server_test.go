package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETStatusOk(t *testing.T) {
	router := SmolMux{m: make(MuxPathEntry)}
	router.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello GET"))
	})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, r)

	want := http.StatusOK
	got := w.Result().StatusCode

	if want != got {
		t.Errorf("want %d got %d", want, got)
	}
}

func TestGETStatusNotFound(t *testing.T) {
	router := SmolMux{m: make(MuxPathEntry)}
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, r)

	want := http.StatusNotFound
	got := w.Result().StatusCode

	if want != got {
		t.Errorf("want %d got %d", want, got)
	}
}

func TestPOSTStatusOk(t *testing.T) {
	router := SmolMux{m: make(MuxPathEntry)}
	router.POST("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello GET"))
	})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodPost, "/", nil)
	router.ServeHTTP(w, r)

	want := http.StatusOK
	got := w.Result().StatusCode

	if want != got {
		t.Errorf("want %d got %d", want, got)
	}
}

func TestPOSTStatusNotFound(t *testing.T) {
	router := SmolMux{m: make(MuxPathEntry)}
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodPost, "/", nil)
	router.ServeHTTP(w, r)

	want := http.StatusNotFound
	got := w.Result().StatusCode

	if want != got {
		t.Errorf("want %d got %d", want, got)
	}
}
