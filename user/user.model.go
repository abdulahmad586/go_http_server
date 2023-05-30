package user

import (
	"encoding/json"
)

type User struct {
	Id       int64  `json:"id"`
	FName    string `json:"fName"`
	LName    string `json:"lName"`
	Age      int8   `json:"age"`
	Password string `json:"password"`
}

func (p User) ToJson() []byte {
	p.Password = "***"
	js, _ := json.Marshal(p)

	return js
}

func FromJson(js []byte) *User {
	p := &User{}
	json.Unmarshal(js, p)

	return p
}
