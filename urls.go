package main

import (
	"github.com/gorilla/mux"
	"WebService/apps/users"
)


//DefineUrls Set general urls for the websevice
func DefineUrls(router *mux.Router){
	// SubRouters
	users.DefineSubUrls(router.PathPrefix("/users").Subrouter())
	
}