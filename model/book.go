package model

import (
	"database/sql"
	"fmt"
)

type BookData struct {
	Id     int
	Isbn   string
	Title  string
	Author string
	Price  float64
}

func GetAllBooks(DB *sql.DB) []BookData {

	var books []BookData
	sql_query := "SELECT * FROM books"

	rows, err := DB.Query(sql_query)

	if err != nil {
		fmt.Println("Error in SQL exec: ", err)
	}

	defer DB.Close()

	for rows.Next() {

		var Id int
		var Isbn string
		var Title string
		var Author string
		var Price float64

		err := rows.Scan(&Id, &Isbn, &Title, &Author, &Price)

		if err != nil {
			panic(err)
		}

		books = append(books, BookData{Id: Id, Isbn: Isbn, Title: Title, Author: Author, Price: Price})
	}

	err = rows.Err() // get any error encountered ing iteration

	if err != nil {
		fmt.Println("Error while fetching books from DB: ", err)
	}
	return books
}
