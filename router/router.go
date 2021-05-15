package router

import (
	"fmt"
	"net/http"
	"os"
)

func IndexHanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page for Rohans website")
}

func RoutesHandler() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	http.HandleFunc("/", IndexHanlder)
}

func BuildServer() {

	err := http.ListenAndServe(":8001", nil)

	if err != nil {
		fmt.Println("Exited due to error", err)
		os.Exit(0)
	}
}
