package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/GoesToEleven/golang-web-dev/042_mongodb/04_controllers/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

// Methods have to be capitalized to be exported, eg, GetUser and not getUser
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    32,
		Id:     p.ByName("id"),
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = "007"

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: only write code to delete user
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Write code to delete user\n")
}
