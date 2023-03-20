package books

import (
	"encoding/json"
	"github.com/gorilla/mux"
	cache "go_intro/utils"
	"io"
	"net/http"
	"time"
)

type Handler struct {
	Dao    Dao
	RCache *cache.Cache
}

func (h *Handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, isExist := h.RCache.Get(r.URL.String())

	if !isExist {
		res, err := h.Dao.GetAllBooks()
		if err != nil {
			w = buildDbBadResponse(w, err)
			return
		}
		h.RCache.Set(r.URL.String(), &res)
		books = res
	}

	data, _ := json.Marshal(books)
	w.Write(data)
}

func (h *Handler) GetBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	books, err := h.Dao.GetBook(id)
	if err != nil {
		w = buildDbBadResponse(w, err)
		return
	}

	data, _ := json.Marshal(books)
	w.Write(data)
}

func (h *Handler) PutBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book Book
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w = buildBadResponse(w, err)
	}

	err = json.Unmarshal(body, &body)
	if err != nil {
		w = buildBadResponse(w, err)
		return
	}

	err = h.Dao.AddNewBook(book)
	if err != nil {
		w = buildDbBadResponse(w, err)
		return
	}

	w.Write(body)
}

func buildDbBadResponse(w http.ResponseWriter, err error) http.ResponseWriter {
	error := Error{
		Timestamp:    time.Now(),
		ErrorMessage: "Something wrong happened during execute SQL query: " + err.Error(),
	}
	r, _ := json.Marshal(error)

	w.WriteHeader(500)
	w.Write(r)
	return w
}

func buildBadResponse(w http.ResponseWriter, err error) http.ResponseWriter {
	error := Error{
		Timestamp:    time.Now(),
		ErrorMessage: "Something wrong happened during request" + err.Error(),
	}
	r, _ := json.Marshal(error)

	w.WriteHeader(500)
	w.Write(r)
	return w
}
