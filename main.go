package main

import (
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("frontend/index.html")

	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}

	data := struct {
		Name string
	}{
		Name: "Go Developer",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		return
	}
}

func staticFileServer(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./frontend"+r.URL.Path)
}

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/static/", staticFileServer)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		return
	}

}
