package controllers

import (
	"encoding/json"
	"first_go_app/models"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type UserController struct {
	userRepo models.UserRepository
}

func NewUserController(userRepo models.UserRepository) *UserController {
	return &UserController{
		userRepo: userRepo,
	}
}

func (controller *UserController) FindById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	user := controller.userRepo.FindByID(id)

	json.NewEncoder(w).Encode(user)
}

func (controller *UserController) Save(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newUser, _ := controller.userRepo.Save(&user)
	json.NewEncoder(w).Encode(newUser)
}

func (controller *UserController) FindAll(w http.ResponseWriter, r *http.Request) {
	allUsers, _ := controller.userRepo.FindAllUsers()
	json.NewEncoder(w).Encode(allUsers)
}

func (controller *UserController) Debit(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var amount struct {
		Amount float64
	}
	err := json.NewDecoder(r.Body).Decode(&amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = controller.userRepo.Debit(id, amount.Amount)
	if err != nil {
		fmt.Println(err)
	}
}

func (controller *UserController) Credit(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var amount struct {
		Amount float64
	}
	err := json.NewDecoder(r.Body).Decode(&amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = controller.userRepo.Credit(id, amount.Amount)
	if err != nil {
		fmt.Println(err)
	}
}

func (controller *UserController) DeleteAcc(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := controller.userRepo.Delete(id)
	if err != nil {
		fmt.Println(err)
	}
}
