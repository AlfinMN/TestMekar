package user

import (
	"BackendTestMekar/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	guuid "github.com/google/uuid"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	DataUsecase UserUsecase
}

func UserController(r *mux.Router, DataUser UserUsecase) {

	handler := UserHandler{DataUser}

	// user
	r.HandleFunc("/users", handler.GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/user", handler.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", handler.DeleteUser).Methods(http.MethodDelete)

	// profession
	r.HandleFunc("/profession", handler.GetProfession).Methods(http.MethodGet)

	// Education
	r.HandleFunc("/education", handler.GetEducation).Methods(http.MethodGet)

}

// get all data user
func (DataUser *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	data, err := DataUser.DataUsecase.GetData()

	if err != nil {
		w.Write([]byte("Data not found"))
	}

	var status model.Response
	status.Status = http.StatusOK
	status.Messages = "data Users Success"
	status.Data = data

	byteOfUser, err := json.Marshal(status)
	if err != nil {
		w.Write([]byte("Oops something wrong "))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfUser)
}

// create user
func (DataUser UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}

	user.UserId = guuid.New().String()

	err = DataUser.DataUsecase.CreateUser(user)
	if err != nil {
		log.Println(err)
	}
	w.Write([]byte("Insert successful"))
}

// get list profession
func (DataUser *UserHandler) GetProfession(w http.ResponseWriter, r *http.Request) {

	data, err := DataUser.DataUsecase.GetProfession()

	if err != nil {
		w.Write([]byte("Data not found"))
	}
	var status model.Response
	status.Status = http.StatusOK
	status.Messages = "data Profession "
	status.Data = data

	byteOfUser, err := json.Marshal(status)
	if err != nil {
		w.Write([]byte("Oops something wrong "))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfUser)

}

func (DataUser *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Println("ini id", id)
	err := DataUser.DataUsecase.DeleteUser(id)
	if err != nil {
		return
	}
	w.Write([]byte("Delete successful"))
}

// get list Education
func (DataUser *UserHandler) GetEducation(w http.ResponseWriter, r *http.Request) {

	data, err := DataUser.DataUsecase.GetEducation()

	if err != nil {
		w.Write([]byte("Data not found"))
	}
	var status model.Response
	status.Status = http.StatusOK
	status.Messages = "data Education "
	status.Data = data

	byteOfUser, err := json.Marshal(status)
	if err != nil {
		w.Write([]byte("Oops something wrong "))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfUser)

}
