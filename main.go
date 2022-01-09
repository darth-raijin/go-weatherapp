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
	data := map[string]interface{}{
		"title": "Base template example",
		"myvar": "Variable example",
	}
	tmpl := template.Must(template.ParseFiles("templates/fragments/base.html"))
	log.Println(tmpl)
	tmpl.Execute(w, data)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Root here")
}

func view_city(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"title": "City - Weatheroo",
		"myvar": "Variable example",
		"city":  "is-primary",
	}
	if r.URL.Query() != nil {
		var city_param = r.URL.Query().Get("city")
		data["title"] = city_param + " - " + "Weatheroo"

	}
	base := template.Must(template.ParseFiles("templates/fragments/base.html"))
	city := template.Must(base.ParseFiles("templates/city.html"))

	log.Println("City GET request")
	city.Execute(w, data)
}

func get_city(city string) {
	// Contact extern API for weatherdata
}

func view_country(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"title":   "Country - Weatheroo",
		"myvar":   "Variable example",
		"country": "is-info",
	}

	base := template.Must(template.ParseFiles("templates/fragments/base.html"))
	city := template.Must(base.ParseFiles("templates/country.html"))

	log.Println("Country GET request")
	city.Execute(w, data)
}

func view_surprise(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"title":    "Country - Weatheroo",
		"myvar":    "Variable example",
		"surprise": "is-info",
	}

	base := template.Must(template.ParseFiles("templates/fragments/base.html"))
	city := template.Must(base.ParseFiles("templates/country.html"))

	log.Println("Country GET request")
	city.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("static")))
	mux.Handle("/css/", twhandler.New(http.Dir("css"), "/css", twembed.New()))

	mux.HandleFunc("/test", test_fragment)
	mux.HandleFunc("/city", view_city)
	mux.HandleFunc("/country", view_country)
	mux.HandleFunc("/surprise", view_surprise)

	http.ListenAndServe(":8080", mux)
	fmt.Println("viewing specific city")

}
