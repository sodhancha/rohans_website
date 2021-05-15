package main

import (
	"fmt"
	"net/http"
	"os"
)

//TODO
/*
* Make a webserver - DONE
* Make a static file handler - DONE
* Make a Routes handler
* Make a DB Connector
* Make a dynamic page which renders SQL output
* Make an admin panel where the dynamic content can be updated from
* Build a RESTFul API for the Backend Content
* Render the API first content in the dynamic web pages
 */

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page for Rohans website")
}

func main() {
	fmt.Println("Starting Rohans Website")

	fileServer := http.FileServer(http.Dir("./static"))

	//not sure what the StripPrefix does. I think it removes the "/static" if static file of that folder is requested let me check
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	//http.Handle("/static/", fileServer) //now checking the same without the strip prefix
	//Without Strip Prefix doesnt seem to serve the static file at all

	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8001", nil)

	if err != nil {
		fmt.Println("Exited due to error", err)
		os.Exit(0)
	}
}
