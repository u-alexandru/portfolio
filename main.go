package main

import (
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

	tmpl.Execute(w, data)
}

func staticFileServer(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./frontend"+r.URL.Path)
}

func main() {
	http.HandleFunc("/static/", staticFileServer)

	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		return
	}

}
