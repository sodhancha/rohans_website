package model

import (
	"database/sql"
	"fmt"
	"os"
)

type Book struct {
	Id     int
	Isbn   string
	Title  string
	Author string
	Price  float64
}

func GetAllBooks(DB *sql.DB) {

	sql_query := "SELECT * FROM books"

	rows, err := DB.Query(sql_query)

	if err != nil {
		fmt.Println("Error in SQL exec: ", err)
	}

	defer DB.Close()

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
