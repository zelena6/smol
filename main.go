package main

import (
	"net/http"
)

func main() {
	c := &SmolMux{m: make(MuxPathEntry)}

	c.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello GET"))
	})

	c.POST("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello POST"))
	})

	http.ListenAndServe(":8080", c)
}
