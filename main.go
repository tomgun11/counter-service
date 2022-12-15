package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var counter = 0

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

	s := "<html><body>Counter: {{.}}</body></html>\n"

	tmpl := template.Must(template.New("test").Parse(s))

	tmpl.Execute(os.Stdout, counter)
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
