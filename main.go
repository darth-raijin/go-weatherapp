package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func test_fragment(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/header.html"))
	tmpl.Execute(w, r)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Root here")
}

func view_city(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "viewing specific city")
}

func main() {
	http.HandleFunc("/", logging(root))
	http.HandleFunc("/city", logging(view_city))
	http.HandleFunc("/t", logging(test_fragment))

	http.ListenAndServe(":8080", nil)
	log.Println("Weather application is good to Go!")
}
