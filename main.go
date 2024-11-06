package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	db   *sql.DB
	once sync.Once
)

type Room struct {
	Id         int       `json:"id"`
	RoomNumber string    `json:"room_number"`
	InProgress int       `json:"in_progress"`
	CreatedAt  time.Time `json:"created_at"`
}

func InitDB() *sql.DB {
	once.Do(func() {
		var err error
		db, err = sql.Open("sqlite3", "./mydatabase.sqlite")
		if err != nil {
			log.Fatal(err)
		}

		// Create tables if they don't exist
		_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS rooms (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				room_number TEXT NOT NULL UNIQUE,
				in_progress INTEGER NOT NULL DEFAULT 1,
				created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
			)
		`)
		if err != nil {
			log.Fatal(err)
		}
	})
	return db
}

func getRoomHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		err := r.ParseForm()
		if err != nil {
			return
		}
		fmt.Println("POST ROOM")
		fmt.Println(r.Form.Get("room_number"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<p>Success</p>"))
	}
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
	http.ServeFile(w, r, "./frontend"+r.URL.Path)
}

func main() {
	InitDB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	http.HandleFunc("/", handler)
	http.HandleFunc("/static/", staticFileServer)
	http.HandleFunc("/rooms/get", getRoomHandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		return
	}

}
