package main

import (
	"compress/gzip"
	"html/template"
	"net/http"
	"strings"
)

func gzipHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the client accepts gzip encoding
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			h(w, r)
			return
		}

		// Create gzip writer and wrap the response writer
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer func(gz *gzip.Writer) {
			err := gz.Close()
			if err != nil {

			}
		}(gz)

		// Use gzipResponseWriter to wrap the original writer
		gzw := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		h(gzw, r)
	}
}

type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (gzw gzipResponseWriter) Write(b []byte) (int, error) {
	return gzw.Writer.Write(b)
}

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
	w.Header().Set("Cache-Control", "public, max-age=31536000")
	w.Header().Set("Expires", "Tue, 01 Jan 2030 00:00:00 GMT")

	http.ServeFile(w, r, "./frontend"+r.URL.Path)
}

func robotsTxtHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	http.ServeFile(w, r, "./frontend/static/robots.txt")
}

func main() {

	http.HandleFunc("/", gzipHandler(handler))
	http.HandleFunc("/static/", gzipHandler(staticFileServer))
	http.HandleFunc("/robots.txt", robotsTxtHandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		return
	}

}
