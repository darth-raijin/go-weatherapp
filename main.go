package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gotailwindcss/tailwind/twembed"
	"github.com/gotailwindcss/tailwind/twhandler"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func test_fragment(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/fragments/header.html"))
	log.Println("yessir")
	tmpl.Execute(w, "s")
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Root here")
}

func view_city(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/city.html"))
	log.Println("yessir")
	tmpl.Execute(w, "s")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("static")))
	mux.Handle("/css/", twhandler.New(http.Dir("css"), "/css", twembed.New()))

	mux.HandleFunc("/test", test_fragment)
	mux.HandleFunc("/city", view_city)

	http.ListenAndServe(":8080", mux)
	fmt.Println("viewing specific city")

}
