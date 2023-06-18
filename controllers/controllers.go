package controllers

import (
	"GOLNGCOURSE/helper"
	"GOLNGCOURSE/utils"
	"encoding/json"
	"net/http"
)

var NewUser helper.UserData

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
