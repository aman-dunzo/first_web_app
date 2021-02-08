package main

import (
	"first_go_app/controllers"
	"first_go_app/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initialMigration() *gorm.DB {
	db, err := gorm.Open("sqlite3", "bank.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&models.User{})

	return db
}

func handleRequests(handler *controllers.UserController) {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/user/{id}", handler.FindById).Methods("GET")
	myRouter.HandleFunc("/user", handler.Save).Methods("POST")
	myRouter.HandleFunc("/users", handler.FindAll).Methods("GET")
	myRouter.HandleFunc("/user/debit/{id}", handler.Debit).Methods("PUT")
	myRouter.HandleFunc("/user/credit/{id}", handler.Credit).Methods("PUT")
	myRouter.HandleFunc("/user/delete/{id}", handler.DeleteAcc).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	db := initialMigration()
	userRepo := models.NewUserRepo(db)
	handler := controllers.NewUserController(userRepo)

	handleRequests(handler)
}
