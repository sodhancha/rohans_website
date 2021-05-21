package router

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/sodhancha/rohans_website/model"
)

type HomePageData struct {
	Title string
	Cats  []CatData
	Books []model.BookData
	Book  model.BookData
}

type CatData struct {
	Name string
}

func CatsCollection() []CatData {
	var cats []CatData
	cats = append(cats, CatData{Name: "Garfiled"})
	cats = append(cats, CatData{Name: "Seinfield"})
	cats = append(cats, CatData{Name: "Cat 3"})
	return cats
}

func IndexHanlder(w http.ResponseWriter, r *http.Request) {

	index_template, err := template.ParseFiles("./templates/home.html")

	if err != nil {
		fmt.Println("Error parsing template: ", err)
	}

	fmt.Println("Staring to parse the template templates/home.html")

	model.GetDBConnection()
	books := model.GetAllBooks(model.DB)

	index_template.Execute(w, HomePageData{Title: "Home Page YO!", Cats: CatsCollection(), Books: books})
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {

	var err error
	var book_id int

	id := r.URL.Query().Get("id")

	book_id, err = strconv.Atoi(id)

	if err != nil {
		fmt.Fprintf(w, "Error parsing request params")
	}

	model.GetDBConnection()
	model.DeleteBookById(book_id, model.DB)

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {

	book_template, err := template.ParseFiles("./templates/book.html")

	if err != nil {
		fmt.Println("Error parsing template: ", err)
	}

	id := r.URL.Query().Get("id")
	model.GetDBConnection()
	book := model.GetBookById(id, model.DB)

	book_template.Execute(w, HomePageData{Title: "Book Edit Page", Book: book})
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {

	var book model.BookData
	var err error

	book.Id, err = strconv.Atoi(r.PostFormValue("id"))

	if err != nil {
		fmt.Println("There was an error parsing Form values", err)
	}

	book.Isbn = r.PostFormValue("isbn")
	book.Title = r.PostFormValue("title")
	book.Price, err = strconv.ParseFloat(r.PostFormValue("price"), 32)
	book.Author = r.PostFormValue("author")

	if err != nil {
		fmt.Println("There was an error parsing Form values", err)
	}

	model.GetDBConnection()
	model.UpdateBookById(book, model.DB)

	edit_id := fmt.Sprint(book.Id)

	http.Redirect(w, r, "/book/edit/?id="+edit_id, http.StatusPermanentRedirect)

}

func RoutesHandler() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	http.HandleFunc("/", IndexHanlder)
	http.HandleFunc("/book/edit/", EditHandler)
	http.HandleFunc("/book/update/", UpdateHandler)
	http.HandleFunc("/book/delete/", DeleteHandler)
}

func BuildServer() {

	err := http.ListenAndServe(":8001", nil)

	if err != nil {
		fmt.Println("Exited due to error", err)
		os.Exit(0)
	}
}
