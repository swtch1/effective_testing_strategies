package main

import "testing"

func TestChangeUserLastName(t *testing.T) {
	db = getRealDB()

	firstName := "Sarah"
	lastName := "Smith"
	newLastName := "Brown"

	// seed the database
	AddUser(&User{
		FName: firstName,
		LName: lastName,
	})

	// update user last name
	ChangeUserLastName("Sarah", "Smith", newLastName)

	// get again with new last name
	updatedUser := GetUserByName(firstName, newLastName)
	if updatedUser.LName != newLastName {
		t.Fatalf("want %v, have %v", newLastName, updatedUser.LName)
	}
}
