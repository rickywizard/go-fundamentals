package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (dd/mm/yyyy): ")

	appAdmin := user.NewAdmin("test@ex.com", "test1234")

	appAdmin.OutputUserDetails()

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	// outputUserDetails(&appUser)
	appUser.OutputUserDetails()

	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

// func outputUserDetails(u *User) {
// 	fmt.Println((*u).firstName, u.lastName, u.birthdate)
// }

func getUserData(prompt string) string {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)

	return value
}