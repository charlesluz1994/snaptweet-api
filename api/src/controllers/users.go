package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		log.Fatal(err)
	}

	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	userId, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(fmt.Sprintf("User created with id: %d", userId)))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Busca 1"))

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Busca todos"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update"))

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete"))
}
