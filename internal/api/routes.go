package api

import (
	"app/internal/models"
	"app/pkg/logs"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (api *APIService) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to the Library Management System API!\n")

	err := logs.LogWriter(r.Method, "/", http.StatusOK)
	if err != nil {
		panic(err)
	}
}

// * Member
func (api *APIService) GetAllMembersHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	members := repository.GetAllMembers()
	fmt.Fprint(w, members)

	err := logs.LogWriter(r.Method, "/members", http.StatusOK)
	if err != nil {
		panic(err)
	}
}

func (api *APIService) GetMemberHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("memberID")
	member := repository.GetMember(id)
	fmt.Fprint(w, member)

	err := logs.LogWriter(r.Method, "/members/{memberID}", http.StatusOK)
	if err != nil {
		panic(err)
	}
}

func (api *APIService) AddMemberHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var newMember models.Member
	err := json.NewDecoder(r.Body).Decode(&newMember)
	if err != nil {
		fmt.Fprintf(w, "Error decoding request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	repository.AddMember(newMember.Name, newMember.Address, newMember.Email)

	fmt.Fprint(w, "Member added successfully!")
	w.WriteHeader(http.StatusCreated)

	err = logs.LogWriter(r.Method, "/members", http.StatusCreated)
	if err != nil {
		panic(err)
	}
}

func (api *APIService) UpdateMemberHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("memberID")

	var newMember models.Member
	err := json.NewDecoder(r.Body).Decode(&newMember)
	if err != nil {
		fmt.Fprintf(w, "Error decoding request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	repository.UpdateMember(id, newMember.Name, newMember.Address, newMember.Email)

	fmt.Fprint(w, "Member updated successfully!")

	err = logs.LogWriter(r.Method, "/members/{memberID}", http.StatusCreated)
	if err != nil {
		panic(err)
	}
}

func (api *APIService) DeleteMemberHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("memberID")
	repository.DeleteMember(id)

	fmt.Fprint(w, "Member deleted successfully!")

	err := logs.LogWriter(r.Method, "/members/{memberID}", http.StatusOK)
	if err != nil {
		panic(err)
	}
}

// * Book
func (api *APIService) GetAllBooksHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	books := repository.GetAllBooks()
	fmt.Fprint(w, books)

	err := logs.LogWriter(r.Method, "/books", http.StatusOK)
	if err != nil {
		panic(err)
	}
}

func (api *APIService) GetBookHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("bookID")
	book := repository.GetBook(id)
	fmt.Fprint(w, book)

	err := logs.LogWriter(r.Method, "/books/{bookID}", http.StatusOK)
	if err != nil {
		panic(err)
	}
}

func (api *APIService) AddBookHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var newBook models.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		fmt.Fprintf(w, "Error decoding request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	repository.AddBook(newBook.Title, newBook.Author, newBook.PublicationYear, newBook.Genre, newBook.TotalCopies)

	fmt.Fprint(w, "Book added successfully!")
	w.WriteHeader(http.StatusCreated)

	err = logs.LogWriter(r.Method, "/books", http.StatusCreated)
	if err != nil {
		panic(err)
	}
}

func (api *APIService) UpdateBookHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("bookID")

	var newBook models.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		fmt.Fprintf(w, "Error decoding request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	repository.UpdateBook(id, newBook.Title, newBook.Author, newBook.PublicationYear, newBook.Genre, newBook.AvailableCopies, newBook.TotalCopies)

	fmt.Fprint(w, "Book updated successfully!")

	err = logs.LogWriter(r.Method, "/books/{bookID}", http.StatusCreated)
	if err != nil {
		panic(err)
	}
}

func (api *APIService) DeleteBookHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("bookID")
	repository.DeleteBook(id)

	fmt.Fprint(w, "Book deleted successfully!")

	err := logs.LogWriter(r.Method, "/books/{bookID}", http.StatusOK)
	if err != nil {
		panic(err)
	}
}
