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
	daysOfWeek := []string{"月", "火", "水", "木", "金", "土", "日"}
	// daysOfWeek := []string{}
	t.Execute(w, daysOfWeek)
}

func checkWith(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("with.html")
	t.Execute(w, "hello")
}

func checkInclude(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t1.html", "t2.html")
	t.Execute(w, "Hello World!")
}

func checkTimeFormat(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatData}
	t := template.New("timeformat.html").Funcs(funcMap)
	t, _ = t.ParseFiles("timeformat.html")
	t.Execute(w, time.Now())
}

func checkContext(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("context.html")
	content := `I asked: <i>"What's up?"</i>`
	t.Execute(w, content)
}

func checkXSS(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("X-XSS-Protection", "0") //chromeとかのxss保護を解除
	t, _ := template.ParseFiles("xss.html")
	t.Execute(w, r.FormValue("comment"))
	// t.Execute(w, template.HTML(r.FormValue("comment"))) //firefoxならalertが反応
}

func formatData(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func checkLayout(w http.ResponseWriter, r *http.Request) {
	// t, _ := template.ParseFiles("layout.html")
	// t.ExecuteTemplate(w, "layout", "")

	rand.Seed(time.Now().Unix())
	var t *template.Template

	// if rand.Intn(10) > 5 {
	// 	t, _ = template.ParseFiles("layout.html", "red_hello.html")
	// } else {
	// 	t, _ = template.ParseFiles("layout.html", "blue_hello.html")
	// }

	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("layout.html", "red_hello.html")
	} else {
		t, _ = template.ParseFiles("layout.html")
	}

	t.ExecuteTemplate(w, "layout", "")
}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("form.html")
	t.Execute(w, nil)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/ifelse", checkIfElse)
	http.HandleFunc("/range", checkRange)
	http.HandleFunc("/with", checkWith)
	http.HandleFunc("/include", checkInclude)
	http.HandleFunc("/time", checkTimeFormat)
	http.HandleFunc("/context", checkContext)
	http.HandleFunc("/xss", checkXSS)
	http.HandleFunc("/form", form)
	http.HandleFunc("/layout", checkLayout)
	server.ListenAndServe()
}
