package users

import (
	userControllers "WebService/apps/users/controllers"
	"github.com/gorilla/mux"
)

// DefineSubUrls Define all users Urls
func DefineSubUrls(s *mux.Router){
	s.HandleFunc("/", userControllers.ListUserController).Methods("GET")
	s.HandleFunc("/create", userControllers.CreateUserController).Methods("POST")
}