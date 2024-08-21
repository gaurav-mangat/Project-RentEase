package ui

import (
	"fmt"
)

func (ui *UI) AppDashboard() {

	// Message to the user
	fmt.Println("**************************************")
	fmt.Println("*                                    *")
	fmt.Println("*           RentEase                 *")
	fmt.Println("*                                    *")
	fmt.Println("**************************************")

	fmt.Println("\n\n 1. LogIn")
	fmt.Println("----------------------")
	fmt.Println(" 2. SignUp")
	fmt.Println("----------------------")
	fmt.Println(" 3. Exit")
	fmt.Println()

	var choice int

	fmt.Print("Enter your choice: ")
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if choice == 1 {
		ui.LoginDashboard()
	} else if choice == 2 {
		ui.SignUpDashboard()
	} else if choice == 3 {
		fmt.Println("Successfully exited the program....")
		return
	} else {
		fmt.Println("\nInvalid choice ")
	}
}
