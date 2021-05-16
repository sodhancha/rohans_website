package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type Book struct {
	Id     int
	Isbn   string
	Title  string
	Author string
	Price  float64
}

type Connect struct {
	KTMDb *sql.DB
}

func GetAllBooks() {

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
		fmt.Println("\n----------------------------")
	}
	err = rows.Err() // get any error encountered ing iteration

	if err != nil {
		os.Exit(0)
	}
}
