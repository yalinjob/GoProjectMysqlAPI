package controllers

import (
	"GOLNGCOURSE/helper"
	"GOLNGCOURSE/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewUser helper.UserData

//var userDetails helper.UserData

func GetUserById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {

		fmt.Println("error while parsing ")

	}
	userDetails, _ := helper.GetUserByID(ID)
	res, _ := json.Marshal(userDetails)

	w.Header().Set("Content-Type", "pkglicaiton/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetAllUser(w http.ResponseWriter, r *http.Request) {

	newUsers := helper.GetAllUser()
	res, _ := json.Marshal(newUsers)
	w.Header().Set("Content-Type", "pkglicaiton/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func CreateUser(w http.ResponseWriter, r *http.Request) {

	CreateUserData := &helper.UserData{}
	utils.ParseBody(r, CreateUserData)
	ud := CreateUserData.CreateUser() // CreateUSerData alreday include the helper package this is the reson we approch diercatly to CreateUser Funcation .
	res, _ := json.Marshal(ud)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {

		fmt.Println("error while parsing ")

	}
	userDe := helper.DeleteUser(ID)
	res, _ := json.Marshal(userDe)

	w.Header().Set("Content-Type", "pkglicaiton/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser = &helper.UserData{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {

		fmt.Println("error while parsing ")

	}
	userDetails, db := helper.GetUserByID(ID)
	if updateUser.FirstName != "" {
		userDetails.FirstName = updateUser.FirstName

	}
	if updateUser.LastName != "" {
		userDetails.LastName = updateUser.LastName

	}
	if updateUser.Email != "" {
		userDetails.Email = updateUser.Email

	}
	if updateUser.NumberOfTicket != 0 {
		userDetails.NumberOfTicket = updateUser.NumberOfTicket

	}
	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)

	w.Header().Set("Content-Type", "pkglicaiton/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
