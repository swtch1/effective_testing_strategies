package main

import (
	"database/sql"
)

func main() {
	db, _ := sql.Open("mysql", "root:password@tcp(localhost:3306)/user")

	udb := &UserDB{
		db: db,
	}

	ChangeUserLastName(udb, "Sarah", "Smith", "Brown")
}

func ChangeUserLastName(udb *UserDB, firstName, lastName, newLastName string) {
	userFromDB := udb.GetUserByName(firstName, lastName)
	userFromDB.LName = newLastName
	udb.UpdateUser(&userFromDB)
}

type User struct {
	ID    int64
	FName string
	LName string
}

type UserDB struct {
	db *sql.DB
}

func (udb *UserDB) AddUser(u *User) {
	q := `INSERT INTO user (first_name, last_name)
				VALUES (?, ?);`
	result, _ := udb.db.Exec(q, u.FName, u.LName)
	u.ID, _ = result.LastInsertId()
}

func (udb *UserDB) UpdateUser(u *User) {
	q := `UPDATE user SET
					first_name = ?,
					last_name = ?
				WHERE id = ?;`
	result, _ := udb.db.Exec(q, u.FName, u.LName, u.ID)
	u.ID, _ = result.LastInsertId()
}

func (udb *UserDB) GetUserByName(fName, lName string) User {
	q := `SELECT
					id,
					first_name,
					last_name
				FROM user
				WHERE first_name = ?
				AND last_name = ?`
	row := udb.db.QueryRow(q, fName, lName)

	var u User
	_ = row.Scan(&u.ID, &u.FName, &u.LName)
	return u
}
