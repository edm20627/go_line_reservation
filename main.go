package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var port int
	if len(os.Args) == 1 {
		port = 8080
	} else {
		port, _ = strconv.Atoi(os.Args[1])
	}
	files := http.FileServer(http.Dir("public"))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", index)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
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
