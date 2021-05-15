package main

import (
	"fmt"
	"net/http"
	"os"
)

//TODO
/*
* Make a webserver
* Make a static file handler
* Make a Routes handler
* Make a DB Connector
* Make a dynamic page which renders SQL output
* Make an admin panel where the dynamic content can be updated from
* Build a RESTFul API for the Backend Content
* Render the API first content in the dynamic web pages
 */

func index(w ResponseWriter, r *Request) {
	fmt.Fprintf(w, "This is the home page for Rohans website")
}

func main() {
	fmt.Println("Starting Rohans Website")

	err := http.ListenAndServe("80801", nil)
	http.HandleFunc("/", index)

	if err != nil {
		fmt.Println("Exited due to error", err)
		os.Exit(0)
	}
}
