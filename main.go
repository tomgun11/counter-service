package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var counter = 0

// Handling function
func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Counter: %v", counter)
	case "POST":
		counter++
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

	s := "<html><body>POST requests that were made: {{.}}</body></html>\n"

	tmpl := template.Must(template.New("test").Parse(s))

	tmpl.Execute(os.Stdout, counter)
}

// Main function that sets the event handler and fires the server
func main() {
	http.HandleFunc("/", hello)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
