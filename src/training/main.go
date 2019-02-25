package main

import (
	"fmt"
	"net/http"
)

type Main struct{}

func (h *Main) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Main")
}

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	main := Main{}
	hello := HelloHandler{}
	world := WorldHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/", &main)
	http.Handle("/hello", &hello)
	http.Handle("/world", &world)
	server.ListenAndServe()
}
