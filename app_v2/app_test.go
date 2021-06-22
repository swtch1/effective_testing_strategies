package main

import "testing"

type UserRepoStub struct {
	user *User
}

func (db *UserRepoStub) GetUserByName(fName, lName string) User {
	return *db.user
}
func (db *UserRepoStub) UpdateUser(u *User) {
	db.user = u
	return
}

func TestChangeUserLastName_Stub(t *testing.T) {
	firstName := "Sarah"
	lastName := "Smith"
	newLastName := "Brown"

	stub := &UserRepoStub{
		user: &User{
			FName: firstName,
			LName: lastName,
		},
	}

	// update user last name
	ChangeUserLastName(stub, firstName, lastName, newLastName)

	// get again with new last name
	updatedUser := stub.GetUserByName(firstName, newLastName)
	if updatedUser.LName != newLastName {
		t.Fatalf("want %v, have %v", newLastName, updatedUser.LName)
	}
}

type UserRepoMock struct {
	GetUserByNameFn func(fName, lName string) User
	UpdateUserFn    func(u *User)
}

func (db *UserRepoMock) GetUserByName(fName, lName string) User {
	return db.GetUserByNameFn(fName, lName)
}
func (db *UserRepoMock) UpdateUser(u *User) {
	db.UpdateUserFn(u)
}

func TestChangeUserLastName_Mock(t *testing.T) {
	firstName := "Sarah"
	lastName := "Smith"
	newLastName := "Brown"

	mock := &UserRepoMock{
		GetUserByNameFn: func(fName, lName string) User {
			return User{
				FName: firstName,
				LName: lastName,
			}
		},
		UpdateUserFn: func(_ *User) {
			lastName = newLastName
		},
	}

	// update user last name
	ChangeUserLastName(mock, firstName, lastName, newLastName)

	// get again with new last name
	updatedUser := mock.GetUserByName(firstName, newLastName)
	if updatedUser.LName != newLastName {
		t.Fatalf("want %v, have %v", newLastName, updatedUser.LName)
	}
}

type UserRepoFake struct {
	c     int64
	users map[string]User
}

func (db *UserRepoFake) AddUser(u *User) {
	db.c++
	u.ID = db.c
	db.users[u.FName+u.LName] = *u
}
func (db *UserRepoFake) GetUserByName(fName, lName string) User {
	return db.users[fName+lName]
}
func (db *UserRepoFake) UpdateUser(u *User) {
	db.users[u.FName+u.LName] = *u
}

func TestChangeUserLastName_Fake(t *testing.T) {
	firstName := "Sarah"
	lastName := "Smith"
	newLastName := "Brown"

	fake := &UserRepoFake{
		users: make(map[string]User),
	}

	fake.AddUser(&User{
		FName: firstName,
		LName: lastName,
	})

	// update user last name
	ChangeUserLastName(fake, firstName, lastName, newLastName)

	// get again with new last name
	updatedUser := fake.GetUserByName(firstName, newLastName)
	if updatedUser.LName != newLastName {
		t.Fatalf("want %v, have %v", newLastName, updatedUser.LName)
	}
}
