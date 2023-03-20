package main

import (
	_ "github.com/lib/pq"
)

const (
	host = ""
	//port     = 5400
	user     = ""
	password = ""
	dbname   = ""
)

func main() {
	//r := mux.NewRouter()
	//dbConn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
	//db, err := sqlx.Connect("postgres", dbConn)
	//defer db.Close()
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//handlers := books.Handler{
	//	RCache: cache.NewCache(),
	//	Dao: books.Dao{
	//		Db: db,
	//	},
	//}
	//booksRouter := r.PathPrefix("/api/books").Subrouter()
	//booksRouter.HandleFunc("", handlers.GetAllBooks).Methods(http.MethodGet)
	//booksRouter.HandleFunc("/{id}", handlers.GetBook).Methods(http.MethodGet)
	//booksRouter.HandleFunc("", handlers.PutBook).Methods(http.MethodPost)
	//
	//log.Fatal(http.ListenAndServe(":8080", r))
}
