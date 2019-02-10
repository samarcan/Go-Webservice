package users

import (
	userRepo "WebService/apps/users/repositories"
	"fmt"
	"encoding/json"
	"net/http"
)

// ListUserController Controller to list users
func ListUserController(w http.ResponseWriter, r *http.Request) {
	result, users := userRepo.ListAllUsers()
	if !result {
		response := `{ "result": false }`
		fmt.Fprintf(w, response)
	} else {
		response, _ := json.Marshal(users)
		fmt.Fprintf(w, string(response))
	}
}
