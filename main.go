package main

import (
	"fmt"
	"log"
	"net/http"
)

// function for handling the hello endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {

	// if given endpoint is not hello endpoint
	if r.URL.Path != "/hello" {
		http.Error(w, "404 : path not found", http.StatusNotFound)
		return
	}

	// if method applied for hello endpoint is wrong
	if r.Method != "GET" {
		http.Error(w, "404 : path not found", http.StatusNotFound)
		return
	}

	// if everything is good then print it
	fmt.Fprint(w, "hello")

}

// function for handling form endpoint
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successfull")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {

	// static files will be served from following directory
	fileServer := http.FileServer(http.Dir("./static"))

	// serve the static files when home is hit by user
	http.Handle("/", fileServer)

	// when /form endpoint is hit then run formHandler function
	http.HandleFunc("/form", formHandler)

	// when /hello endpoint is hit then run helloHandler function
	http.HandleFunc("/hello", helloHandler)

	// indicate the start of server
	fmt.Printf("Starting server at port 3000")

	// start the server to listen and server and if there is error then print it
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
