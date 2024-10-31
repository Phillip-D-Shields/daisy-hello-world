// main.go
package main

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Title    string
	Headline string
	Subtext  string
}

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle root route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		data := PageData{
			Title:    "Demo",
			Headline: "Kia ora!",
			Subtext:  "Go templates with the beauty of DaisyUI",
		}
		tmpl.Execute(w, data)
	})

	println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}
