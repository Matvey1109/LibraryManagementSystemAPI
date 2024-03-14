package api

import (
	"app/internal/serializers"
	"app/pkg/logs"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (api *APIService) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Welcome to the Library Management System API!"))

	err := logs.LogWriter(r.Method, "/", http.StatusOK)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// * Member
func (api *APIService) GetAllMembersHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	members, err := repository.GetAllMembers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logs.LogWriter(r.Method, "/members", http.StatusInternalServerError)
		return
	}

	jsonData, err := serializers.SerializeJsonData(members)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.LogWriter(r.Method, "/members", http.StatusBadRequest)
		return
	}

	w.Write(jsonData)

	err = logs.LogWriter(r.Method, "/members", http.StatusOK)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (api *APIService) GetMemberHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("memberID")
	member, err := repository.GetMember(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		logs.LogWriter(r.Method, "/members/{memberID}", http.StatusNotFound)
		return
	}

	jsonData, err := serializers.SerializeJsonData(member)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.LogWriter(r.Method, "/members/{memberID}", http.StatusBadRequest)
		return
	}

	w.Write(jsonData)

	err = logs.LogWriter(r.Method, "/members/{memberID}", http.StatusOK)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (api *APIService) AddMemberHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := serializers.DeserializeJsonData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.LogWriter(r.Method, "/members", http.StatusBadRequest)
		return
	}

	var (
		name, address, email string
	)

	name, address, email, err = serializers.ValidateAddMemberData(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.LogWriter(r.Method, "/members", http.StatusBadRequest)
		return
	}

	err = repository.AddMember(name, address, email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logs.LogWriter(r.Method, "/members", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Member added successfully!"))

	err = logs.LogWriter(r.Method, "/members", http.StatusCreated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (api *APIService) UpdateMemberHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("memberID")

	data, err := serializers.DeserializeJsonData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.LogWriter(r.Method, "/members/{memberID}", http.StatusBadRequest)
		return
	}

	var (
		name, address, email *string
	)

	name, address, email, err = serializers.ValidateUpdateMemberData(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.LogWriter(r.Method, "/members/{memberID}", http.StatusBadRequest)
		return
	}

	err = repository.UpdateMember(id, name, address, email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		logs.LogWriter(r.Method, "/members/{memberID}", http.StatusNotFound)
		return
	}

	w.Write([]byte("Member updated successfully!"))

	err = logs.LogWriter(r.Method, "/members/{memberID}", http.StatusOK)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (api *APIService) DeleteMemberHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("memberID")
	err := repository.DeleteMember(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		logs.LogWriter(r.Method, "/members/{memberID}", http.StatusNotFound)
		return
	}

	w.Write([]byte("Member deleted successfully!"))

	err = logs.LogWriter(r.Method, "/members/{memberID}", http.StatusOK)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// * Book
func (api *APIService) GetAllBooksHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	books, err := repository.GetAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logs.LogWriter(r.Method, "/books", http.StatusInternalServerError)
		return
	}

	jsonData, err := serializers.SerializeJsonData(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.LogWriter(r.Method, "/books", http.StatusBadRequest)
		return
	}

	w.Write(jsonData)

	err = logs.LogWriter(r.Method, "/books", http.StatusOK)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (api *APIService) GetBookHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("bookID")
	book, err := repository.GetBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		logs.LogWriter(r.Method, "/books/{bookID}", http.StatusNotFound)
		return
	}

	jsonData, err := serializers.SerializeJsonData(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.LogWriter(r.Method, "/books/{bookID}", http.StatusBadRequest)
		return
	}

	w.Write(jsonData)

	err = logs.LogWriter(r.Method, "/books/{bookID}", http.StatusOK)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (api *APIService) AddBookHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := serializers.DeserializeJsonData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.LogWriter(r.Method, "/books", http.StatusBadRequest)
		return
	}

	var (
		title, author, genre         string
		publicationYear, totalCopies int
	)

	title, author, publicationYear, genre, totalCopies, err = serializers.ValidateAddBookData(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.LogWriter(r.Method, "/books", http.StatusBadRequest)
		return
	}

	err = repository.AddBook(title, author, publicationYear, genre, totalCopies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.LogWriter(r.Method, "/books", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Book added successfully!"))

	err = logs.LogWriter(r.Method, "/books", http.StatusCreated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (api *APIService) UpdateBookHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("bookID")

	data, err := serializers.DeserializeJsonData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.LogWriter(r.Method, "/books/{bookID}", http.StatusBadRequest)
		return
	}

	var (
		title, author, genre                          *string
		publicationYear, availableCopies, totalCopies *int
	)

	title, author, publicationYear, genre, availableCopies, totalCopies, err = serializers.ValidateUpdateBookData(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.LogWriter(r.Method, "/books/{bookID}", http.StatusBadRequest)
		return
	}

	err = repository.UpdateBook(id, title, author, publicationYear, genre, availableCopies, totalCopies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.LogWriter(r.Method, "/books/{bookID}", http.StatusBadRequest)
		return
	}

	w.Write([]byte("Book updated successfully!"))

	err = logs.LogWriter(r.Method, "/books/{bookID}", http.StatusOK)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (api *APIService) DeleteBookHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("bookID")
	err := repository.DeleteBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		logs.LogWriter(r.Method, "/books/{bookID}", http.StatusNotFound)
		return
	}

	w.Write([]byte("Book deleted successfully!"))

	err = logs.LogWriter(r.Method, "/books/{bookID}", http.StatusOK)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// * Borrowing
func (api *APIService) GetAllBorrowingsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	borrowings, err := repository.GetAllBorrowings()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logs.LogWriter(r.Method, "/borrowings", http.StatusInternalServerError)
		return
	}

	jsonData, err := serializers.SerializeJsonData(borrowings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.LogWriter(r.Method, "/borrowings", http.StatusBadRequest)
		return
	}

	w.Write(jsonData)

	err = logs.LogWriter(r.Method, "/borrowings", http.StatusOK)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (api *APIService) GetMemberBooksHanlder(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("memberID")
	books, err := repository.GetMemberBooks(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.LogWriter(r.Method, "/borrowings/{memberID}", http.StatusBadRequest)
		return
	}

	jsonData, err := serializers.SerializeJsonData(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.LogWriter(r.Method, "/borrowings/{memberID}", http.StatusBadRequest)
		return
	}

	w.Write(jsonData)

	err = logs.LogWriter(r.Method, "/borrowings/{memberID}", http.StatusOK)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (api *APIService) BorrowBookHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := serializers.DeserializeJsonData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.LogWriter(r.Method, "/borrowings", http.StatusBadRequest)
		return
	}

	var (
		bookID, memberID string
		borrowYear       int
	)

	bookID, memberID, borrowYear, err = serializers.ValidateAddBorrowingData(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.LogWriter(r.Method, "/borrowings", http.StatusBadRequest)
		return
	}

	err = repository.BorrowBook(bookID, memberID, borrowYear)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.LogWriter(r.Method, "/borrowings", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Book borrowed successfully!"))

	err = logs.LogWriter(r.Method, "/borrowings", http.StatusCreated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (api *APIService) ReturnBookHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("borrowingID")
	err := repository.ReturnBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		logs.LogWriter(r.Method, "/borrowings/{borrowingID}", http.StatusNotFound)
		return
	}

	w.Write([]byte("Book returned successfully!"))

	err = logs.LogWriter(r.Method, "/borrowings/{borrowingID}", http.StatusOK)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
