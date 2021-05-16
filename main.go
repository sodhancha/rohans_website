package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sodhancha/rohans_website/model"
	"github.com/sodhancha/rohans_website/router"
)

//TODO
/*
* Make a webserver - DONE
* Make a static file handler - DONE
* Make a Routes handler - DONE
* Make a DB Connector
* Make a dynamic page which renders SQL output
* Make an admin panel where the dynamic content can be updated from
* Build a RESTFul API for the Backend Content
* Render the API first content in the dynamic web pages
 */

func main() {
	fmt.Println("Starting Rohans Website")

	model.GetAllBooks()
	//Remember that Go Lang requires exported functions to have Capital Case Names
	router.RoutesHandler()
	router.BuildServer()

}
