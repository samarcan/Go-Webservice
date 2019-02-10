package users

import (
	userEntities "WebService/apps/users/entities"
	userRepo "WebService/apps/users/repositories"
	"fmt"
	"encoding/json"
	"net/http"
)

// CreateUserController Endpoint to create new users
func CreateUserController(w http.ResponseWriter, r *http.Request) {
	user := &userEntities.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		fmt.Println("Error while decoding request body: ", err)
	}
	result := userRepo.CreateUser(user)

	dataResponse := map[string]bool {
		"result": result,
	}
	response, _ := json.Marshal(dataResponse)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
