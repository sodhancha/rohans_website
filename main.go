package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
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

func DBConnector() {
	os.Remove("./foo.db")

	db, err := sql.Open("sqlite3", "./DB/foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func main() {
	fmt.Println("Starting Rohans Website")
	DBConnector()
	//Remember that Go Lang requires exported functions to have Capital Case Names
	router.RoutesHandler()
	router.BuildServer()

}
