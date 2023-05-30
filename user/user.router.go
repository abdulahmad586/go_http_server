package user

import (
	"encoding/json"
	"fmt"
	"go_http_server/response"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllUsersRoute(w http.ResponseWriter, r *http.Request) {
	response := response.Response{
		Ok:      true,
		Payload: GetAllUsers(),
	}
	w.Write(response.ToJson())
}

func GetOneUserRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	user, err := GetOneUser(id)

	response := response.Response{
		Ok:      err == nil,
		Payload: user,
		Message: fmt.Sprintf("%v", err),
	}

	w.Write(response.ToJson())
}

func AddUserRoute(w http.ResponseWriter, r *http.Request) {

	newUser := User{}

	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	json.Unmarshal(body, &newUser)

	res := AddUser(newUser.FName, newUser.LName, newUser.Password, newUser.Age)

	response := response.Response{
		Ok:      true,
		Payload: res,
	}

	w.Write(response.ToJson())
}

func UpdateUserRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	update := User{}

	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	json.Unmarshal(body, &update)

	res, err := EditUser(id, update.FName, update.LName, update.Age)

	response := response.Response{
		Ok:      err == nil,
		Message: fmt.Sprintf("%v", err),
		Payload: res,
	}

	w.Write(response.ToJson())
}

func DeleteUserRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	res, err := DeleteUser(id)

	response := response.Response{
		Ok:      err == nil,
		Message: fmt.Sprintf("%v", err),
		Payload: res,
	}

	w.Write(response.ToJson())
}
