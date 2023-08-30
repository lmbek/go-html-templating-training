package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Open and read the template file
		file, err := os.Open("index.html")
		if err != nil {
			fmt.Println("1")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Parse the template from the file
		tmpl, err := template.New("index.html").ParseFiles("index.html")
		if err != nil {
			fmt.Println("2")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Define data for the template
		data := struct {
			Title string
			Name  string
		}{
			Title: "Go Templating Example",
			Name:  "John",
		}

		// Execute the template with the provided data
		err = tmpl.Execute(w, data)
		if err != nil {
			fmt.Println("Error executing template:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

	})

	http.ListenAndServe(":8082", nil)
}
