package books

import "time"

type Book struct {
	Id          int64     `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Author      string    `db:"author" json:"author"`
	PublishDate time.Time `db:"publish_date" json:"publishDate"`
}

type Error struct {
	Timestamp    time.Time `json:"timestamp"`
	ErrorMessage string    `json:"errorMessage"`
}
