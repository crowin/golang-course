package books

import "github.com/jmoiron/sqlx"

type Dao struct {
	Db *sqlx.DB
}

func (d *Dao) GetAllBooks() ([]Book, error) {
	var books []Book
	err := d.Db.Select(&books, "SELECT * from books")
	return books, err
}

func (d *Dao) GetBook(id string) (Book, error) {
	var book Book
	err := d.Db.Get(&book, "SELECT * from books WHERE id = $1", id)
	return book, err
}

func (d *Dao) AddNewBook(book Book) error {
	q := "INSERT INTO books (title, author) VALUES (:title, :author)"
	_, err := d.Db.Queryx(q, &book)
	return err
}
