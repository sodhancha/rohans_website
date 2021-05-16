package main

import (
	"database/sql"
	"fmt"
	"log"

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

type Book struct {
	Id     int
	Isbn   string
	Title  string
	Author string
	Price  float64
}

func DBConnector() {

	db, err := sql.Open("sqlite3", "./DB/foo.db")
	if err != nil {
		log.Fatal(err)
	}

	sql_query := "SELECT * FROM books"

	rows, err := db.Query(sql_query)

	if err != nil {
		fmt.Println("Error in SQL exec: ", err)
	}

	defer db.Close()

	for rows.Next() {

		book := Book{}
		err = rows.Scan(&book.Id, &book.Isbn, &book.Title, &book.Author, &book.Price)

		if err != nil {
			panic(err)
		}
		fmt.Print(book.Id)
		fmt.Print(book.Isbn)
		fmt.Print(book.Title)
		fmt.Print(book.Author)
		fmt.Print(book.Price)
		fmt.Println("\n----------------------------\n")
	}
	err = rows.Err() // get any error encountered ing iteration
}

func main() {
	fmt.Println("Starting Rohans Website")
	DBConnector()
	//Remember that Go Lang requires exported functions to have Capital Case Names
	router.RoutesHandler()
	router.BuildServer()

}
