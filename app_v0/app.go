package main

import (
	"database/sql"
)

var db *sql.DB

func main() {
	db, _ = sql.Open("mysql", "root:password@tcp(localhost:3306)/user")

	ChangeUserLastName("Sarah", "Smith", "Brown")
}

func ChangeUserLastName(firstName, lastName, newLastName string) {
	userFromDB := GetUserByName(firstName, lastName)
	userFromDB.LName = newLastName
	UpdateUser(&userFromDB)
}

type User struct {
	ID    int64
	FName string
	LName string
}

func UpdateUser(u *User) {
	q := `UPDATE user SET
					first_name = ?,
					last_name = ?
				WHERE id = ?;`
	result, _ := db.Exec(q, u.FName, u.LName, u.ID)
	u.ID, _ = result.LastInsertId()
}

func GetUserByName(fName, lName string) User {
	q := `SELECT
					id,
					first_name,
					last_name
				FROM user
				WHERE first_name = ?
				AND last_name = ?`
	row := db.QueryRow(q, fName, lName)

	var u User
	_ = row.Scan(&u.ID, &u.FName, &u.LName)
	return u
}
