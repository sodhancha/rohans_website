package model

import (
	"database/sql"
	"fmt"
)

var Id int
var Isbn string
var Title string
var Author string
var Price float64

type BookData struct {
	Id     int
	Isbn   string
	Title  string
	Author string
	Price  float64
}

func InsertNewBook(book BookData, DB *sql.DB) {

	sql_stmt := "UPDATE books SET isbn=?, title=?, author=?, price=? WHERE id=?"

	rows, err := DB.Query(sql_stmt, book.Isbn, book.Title, book.Author, book.Price, book.Id)

	if err != nil {
		fmt.Println("Error in updating SQL: ", err)
	}

	for rows.Next() {
		rows.Scan()
	}

}

func UpdateBookById(book BookData, DB *sql.DB) {

	sql_stmt := "UPDATE books SET isbn=?, title=?, author=?, price=? WHERE id=?"

	rows, err := DB.Query(sql_stmt, book.Isbn, book.Title, book.Author, book.Price, book.Id)

	if err != nil {
		fmt.Println("Error in updating SQL: ", err)
	}

	for rows.Next() {
		rows.Scan()
	}

}

func GetBookById(id string, DB *sql.DB) BookData {

	fmt.Println("ID of the book: ", id)

	row := DB.QueryRow("SELECT * FROM books WHERE id=?", id)

	err := row.Scan(&Id, &Isbn, &Title, &Author, &Price)

	if err != nil {
		fmt.Println("DB Error: ", err)
	}

	book := BookData{Id: Id, Isbn: Isbn, Title: Title, Author: Author, Price: Price}

	return book

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

func DeleteBookById(id int, DB *sql.DB) {

	is_deleted, _ := DB.Exec("DELETE FROM books WHERE id=?", id)

	_, err := is_deleted.RowsAffected()

	if err != nil {
		fmt.Println("Error deleting row in Books: ", err)
	}
}
