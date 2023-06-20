package routes

import (
	"GOLNGCOURSE/controllers"

	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {

	router.HandleFunc("/user/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/", controllers.GetAllUser).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{userId}", controllers.UpdateUser).Methods("PUT")

}
