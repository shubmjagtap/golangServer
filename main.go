package main

import (
	"net/http"
)

func main() {

	// static files will be served from following directory
	fileServer := http.FileServer(http.Dir("./static"))

	// serve the static files when home is hit by user
	http.Handle("/", fileServer)

	// when /form endpoint is hit then run formHandler function
	http.HandleFunc("/form", formHandler)

	// when /hello endpoint is hit then run helloHandler function
	http.HandlerFunc("/hello", helloHandler)

	// code to be written

}
