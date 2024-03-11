package api

import (
	"app/internal/apierrors"
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
		apierrors.PrintError(w, err, http.StatusInternalServerError)
	}
}

// * Member
func (api *APIService) GetAllMembersHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	members, err := repository.GetAllMembers()
	if err != nil {
		apierrors.PrintError(w, err, http.StatusInternalServerError)
		logs.LogWriter(r.Method, "/members", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, members)

	err = logs.LogWriter(r.Method, "/members", http.StatusOK)
	if err != nil {
		apierrors.PrintError(w, err, http.StatusInternalServerError)
	}
}

func (api *APIService) GetMemberHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("memberID")
	member, err := repository.GetMember(id)
	if err != nil {
		apierrors.PrintError(w, err, http.StatusNotFound)
		logs.LogWriter(r.Method, "/members/{memberID}", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, member)

	err = logs.LogWriter(r.Method, "/members/{memberID}", http.StatusOK)
	if err != nil {
		apierrors.PrintError(w, err, http.StatusInternalServerError)
	}
}

func (api *APIService) AddMemberHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var newMember models.Member
	err := json.NewDecoder(r.Body).Decode(&newMember)
	if err != nil {
		apierrors.PrintError(w, err, http.StatusBadRequest)
		logs.LogWriter(r.Method, "/members", http.StatusBadRequest)
		return
	}

	err = repository.AddMember(newMember.Name, newMember.Address, newMember.Email)
	if err != nil {
		apierrors.PrintError(w, err, http.StatusInternalServerError)
		logs.LogWriter(r.Method, "/members", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Member added successfully!")

	err = logs.LogWriter(r.Method, "/members", http.StatusCreated)
	if err != nil {
		apierrors.PrintError(w, err, http.StatusInternalServerError)
	}
}

func (api *APIService) UpdateMemberHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("memberID")

	var newMember models.Member
	err := json.NewDecoder(r.Body).Decode(&newMember)
	if err != nil {
		apierrors.PrintError(w, err, http.StatusBadRequest)
		logs.LogWriter(r.Method, "/members/{memberID}", http.StatusBadRequest)
		return
	}

	err = repository.UpdateMember(id, newMember.Name, newMember.Address, newMember.Email)
	if err != nil {
		apierrors.PrintError(w, err, http.StatusNotFound)
		logs.LogWriter(r.Method, "/members/{memberID}", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "Member updated successfully!")

	err = logs.LogWriter(r.Method, "/members/{memberID}", http.StatusOK)
	if err != nil {
		apierrors.PrintError(w, err, http.StatusInternalServerError)
	}
}

func (api *APIService) DeleteMemberHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("memberID")
	err := repository.DeleteMember(id)
	if err != nil {
		apierrors.PrintError(w, err, http.StatusNotFound)
		logs.LogWriter(r.Method, "/members/{memberID}", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "Member deleted successfully!")

	err = logs.LogWriter(r.Method, "/members/{memberID}", http.StatusOK)
	if err != nil {
		apierrors.PrintError(w, err, http.StatusInternalServerError)
	}
}

// * Book
func (api *APIService) GetAllBooksHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	books, err := repository.GetAllBooks()
	if err != nil {
		apierrors.PrintError(w, err, http.StatusInternalServerError)
		logs.LogWriter(r.Method, "/books", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, books)

	err = logs.LogWriter(r.Method, "/books", http.StatusOK)
	if err != nil {
		apierrors.PrintError(w, err, http.StatusInternalServerError)
	}
}

func (api *APIService) GetBookHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("bookID")
	book, err := repository.GetBook(id)
	if err != nil {
		apierrors.PrintError(w, err, http.StatusNotFound)
		logs.LogWriter(r.Method, "/books/{bookID}", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, book)

	err = logs.LogWriter(r.Method, "/books/{bookID}", http.StatusOK)
	if err != nil {
		apierrors.PrintError(w, err, http.StatusInternalServerError)
	}
}

func (api *APIService) AddBookHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var newBook models.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		apierrors.PrintError(w, err, http.StatusBadRequest)
		logs.LogWriter(r.Method, "/books", http.StatusBadRequest)
		return
	}

	err = repository.AddBook(newBook.Title, newBook.Author, newBook.PublicationYear, newBook.Genre, newBook.TotalCopies)
	if err != nil {
		apierrors.PrintError(w, err, http.StatusInternalServerError)
		logs.LogWriter(r.Method, "/books", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Book added successfully!")

	err = logs.LogWriter(r.Method, "/books", http.StatusCreated)
	if err != nil {
		apierrors.PrintError(w, err, http.StatusInternalServerError)
	}
}

func (api *APIService) UpdateBookHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("bookID")

	var newBook models.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		apierrors.PrintError(w, err, http.StatusBadRequest)
		logs.LogWriter(r.Method, "/books/{bookID}", http.StatusBadRequest)
		return
	}

	err = repository.UpdateBook(id, newBook.Title, newBook.Author, newBook.PublicationYear, newBook.Genre, newBook.AvailableCopies, newBook.TotalCopies)
	if err != nil {
		apierrors.PrintError(w, err, http.StatusNotFound)
		logs.LogWriter(r.Method, "/books/{bookID}", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "Book updated successfully!")

	err = logs.LogWriter(r.Method, "/books/{bookID}", http.StatusOK)
	if err != nil {
		apierrors.PrintError(w, err, http.StatusInternalServerError)
	}
}

func (api *APIService) DeleteBookHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("bookID")
	err := repository.DeleteBook(id)
	if err != nil {
		apierrors.PrintError(w, err, http.StatusNotFound)
		logs.LogWriter(r.Method, "/books/{bookID}", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "Book deleted successfully!")

	err = logs.LogWriter(r.Method, "/books/{bookID}", http.StatusOK)
	if err != nil {
		apierrors.PrintError(w, err, http.StatusInternalServerError)
	}
}
