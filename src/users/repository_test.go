package users

import (
	"ilyaryabchinski/gotask/src/db"
	"os"
	"testing"
)

func prepareData() {
	db.ConnectToDatabase("../../data-test")
	testData := []User{
		{FirstName: "John", LastName: "Doe"},
		{FirstName: "Ada", LastName: "Lovelace"},
	}
	Create(testData[0])
	Create(testData[1])
}

func TestGetAll(t *testing.T) {
	prepareData()

	users, _ := GetAll()

	failure :=
		users[0].FirstName != "John" || users[0].LastName != "Doe" ||
			users[1].FirstName != "Ada" || users[1].LastName != "Lovelace"

	if failure {
		t.Error("Get All failed")
	}

	t.Cleanup(func() {
		db.CloseDatabaseConnection()
		os.RemoveAll("../../data-test")
	})
}

func TestGetOne(t *testing.T) {
	prepareData()

	users, _ := GetAll()

	user, _ := GetOne(users[0].PersonalCode)

	failure := user.FirstName != "John" || user.LastName != "Doe"

	if failure {
		t.Error("Get One failed")
	}

	t.Cleanup(func() {
		db.CloseDatabaseConnection()
		os.RemoveAll("../../data-test")
	})
}

func TestDelete(t *testing.T) {
	prepareData()

	users, _ := GetAll()

	Delete(users[0].PersonalCode)

	usersUpdated, _ := GetAll()

	failure := len(usersUpdated) > 1

	if failure {
		t.Error("Get One failed")
	}

	t.Cleanup(func() {
		db.CloseDatabaseConnection()
		os.RemoveAll("../../data-test")
	})
}

func TestEdit(t *testing.T) {
	prepareData()

	users, _ := GetAll()

	updatedUser := users[0]

	updatedUser.FirstName = "Leo"
	updatedUser.LastName = "Messi"

	Edit(updatedUser.PersonalCode, updatedUser)

	userToCheck, _ := GetOne(updatedUser.PersonalCode)

	failure := userToCheck.FirstName != "Leo" || userToCheck.LastName != "Messi"

	if failure {
		t.Error("Get One failed")
	}

	t.Cleanup(func() {
		db.CloseDatabaseConnection()
		os.RemoveAll("../../data-test")
	})
}
