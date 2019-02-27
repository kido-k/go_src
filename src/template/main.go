package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func checkIfElse(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("if_else.html")
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func checkRange(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("range.html")
	// daysOfWeek := []string{"月", "火", "水", "木", "金", "土", "日"}
	daysOfWeek := []string{}
	t.Execute(w, daysOfWeek)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/ifelse", checkIfElse)
	http.HandleFunc("/range", checkRange)
	server.ListenAndServe()
}
