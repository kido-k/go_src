package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HelloHandler struct{}

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	// h := r.Header["Accept-Encoding"]
	fmt.Fprintln(w, h)
}

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
	// fmt.Fprintln(w, "body")
}

func process(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()           //=>map[hello:[kido_k world] post:[456] thread:[123]]
	// fmt.Fprintln(w, r.Form) //queryデータとpostデータをリストとして渡す
	// fmt.Fprintln(w, r.PostForm) //postデータのみ

	r.ParseMultipartForm(1024) // =>
	// fmt.Fprintln(w, r.MultipartForm) //queryデータとpostデータをリストとして渡す

	fileHeader := r.MultipartForm.File["uploaded"][0]

	file, err := fileHeader.Open()

	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
	fmt.Println("receive")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	// http.HandleFunc("/headers", headers)
	// http.HandleFunc("/body", body)
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
