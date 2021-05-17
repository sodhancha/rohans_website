package router

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type HomePageData struct {
	Title string
}

func IndexHanlder(w http.ResponseWriter, r *http.Request) {

	index_template, err := template.ParseFiles("./templates/home.html")

	fmt.Println("Staring to parse the template templates/home.html")
	if err != nil {
		fmt.Println("Error parsing template: ", err)
	}

	index_template.Execute(w, HomePageData{Title: "Rohand Home Page"})
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
