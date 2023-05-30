package user

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var myDb *sql.DB

func init() {
	var err error
	myDb, err = sql.Open("sqlite3", "mydb.db")
	checkError(err)
	statement, err := myDb.Prepare(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		fName TEXT,
		lName TEXT,
		password TEXT,
		age INTEGER
	);`)
	checkError(err)
	statement.Exec()

}

func GetAllUsers() []*User {

	users := []*User{}

	result, err := myDb.Query("SELECT * FROM users")
	checkError(err)
	for result.Next() {
		var id int64
		var age int8
		var fName, lName, password string
		result.Scan(&id, &fName, &lName, &password, &age)

		users = append(users, &User{Id: id, FName: fName, LName: lName, Password: password, Age: age})
	}
	return users
}

func GetOneUser(i int) (*User, error) {

	result := myDb.QueryRow("SELECT * FROM users WHERE id=?", i)

	if result == nil {
		return nil, errors.New("User not found")
	}

	var id int64
	var age int8
	var fName, lName, password string
	result.Scan(&id, &fName, &lName, &password, &age)

	return &User{Id: id, FName: fName, LName: lName, Password: password, Age: age}, nil
}

func AddUser(fname, lname, password string, age int8) *User {
	user := &User{
		FName:    fname,
		LName:    lname,
		Password: password,
		Age:      age,
	}

	statement, err := myDb.Exec("INSERT INTO users (fName, lName, password, age) VALUES (?,?,?,?)", fname, lname, password, age)
	checkError(err)
	user.Id, err = statement.LastInsertId()
	checkError(err)

	return user
}

func EditUser(id int, fname, lname string, age int8) (*User, error) {
	u, err := GetOneUser(id)
	if err != nil {
		return nil, err
	}
	user := &User{
		Id:       u.Id,
		FName:    fname,
		LName:    lname,
		Password: u.Password,
		Age:      age,
	}

	_, err1 := myDb.Query("UPDATE users SET fName=?, lName=?, age=? where id = ?", fname, lname, age, id)
	if err1 != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(id int) (*User, error) {
	user, err := GetOneUser(id)
	if err != nil {
		return nil, err
	}
	res, err1 := myDb.Exec("DELETE FROM users WHERE id = ?", id)
	if err1 != nil {
		return nil, err
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		return nil, errors.New("Unable to delete item")
	}

	return user, nil
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("An error occurred:%v", err)
	}
}
