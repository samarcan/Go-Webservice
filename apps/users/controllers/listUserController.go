package users

import (
	userRepo "WebService/apps/users/repositories"
	"encoding/json"
	"net/http"
)

// ListUserController Controller to list users
func ListUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, users := userRepo.ListAllUsers()
	if !result {
		response := `{ "result": false }`
		w.Write([]byte(response))
	} else {
		response, _ := json.Marshal(users)
		w.Write(response)
	}
}
