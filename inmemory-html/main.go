package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Define the template text
		templateText := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>{{.Title}}</title>
		</head>
		<body>
			<h1>Hello, {{.Name}}!</h1>
			<p>This is a simple Go HTML templating example.</p>
		</body>
		</html>`

		// Create a new template and parse the template text
		tmpl, err := template.New("index").Parse(templateText)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Define data for the template
		data := struct {
			Title string
			Name  string
		}{
			Title: "Go Templating Example",
			Name:  "Lars",
		}

		// Execute the template with the provided data
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":8082", nil)
}
