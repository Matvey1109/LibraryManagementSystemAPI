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

func (api *APIService) GetAllMembersHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	members := repository.GetAllMembers()
	fmt.Fprint(w, members)

	err := logs.LogWriter(r.Method, "/members", http.StatusOK)
	if err != nil {
		panic(err)
	}
}

func (api *APIService) GetMemberHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	member := repository.GetMember(id)
	fmt.Fprint(w, member)

	err := logs.LogWriter(r.Method, "/members/{id}", http.StatusOK)
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
