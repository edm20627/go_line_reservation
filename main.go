package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	files := http.FileServer(http.Dir("public"))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", index)
	if port == "" {
		http.ListenAndServe(":8080", nil)
	} else {
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	if err := t.ExecuteTemplate(w, "index.html", nil); err != nil {
		log.Fatal(err)
	}
}
