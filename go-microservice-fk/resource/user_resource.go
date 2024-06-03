package resource

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hrishikeshgoyal/quest/go-microservice-fk/entity"
	"github.com/hrishikeshgoyal/quest/go-microservice-fk/service"
	"log"
	"net/http"
)

type UserResource struct {
	userService service.UserService
}

func RegisterUserResource(r *mux.Router, us service.UserService) {
	ur := &UserResource{
		userService: us,
	}
	r.HandleFunc("/users", ur.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", ur.GetUser).Methods("GET")
	r.HandleFunc("/usersn/{name}", ur.GetUserByName).Methods("GET")
	r.HandleFunc("/users", ur.CreateUser).Methods("POST")
}

func (ur *UserResource) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := ur.userService.GetAllUsers()
	json.NewEncoder(w).Encode(users)
}

func (ur *UserResource) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	user := ur.userService.GetUser(params["id"])
	json.NewEncoder(w).Encode(user)
}

func (ur *UserResource) GetUserByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	log.Printf("param{name} = %s", params["name"])
	user := ur.userService.GetUserByName(params["name"])
	json.NewEncoder(w).Encode(user)
}

func (ur *UserResource) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entity.User

	json.NewDecoder(r.Body).Decode(&user)
	log.Printf("Extracted user %v", user)
	createdUser := ur.userService.CreateUser(&user)
	log.Printf("Created user %v", createdUser)
	json.NewEncoder(w).Encode(createdUser)
}
