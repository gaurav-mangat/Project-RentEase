package ui

import (
	"fmt"
	"rentease/pkg/utils"
)

func (ui *UI) sellerDashboard() {
	for {
		// Display the seller dashboard
		fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m") // Sky blue
		fmt.Println("\033[1;31m                        SELLER DASHBOARD                        \033[0m") // Red bold
		fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m") // Sky blue

		fmt.Println("\033[1;32m1. List Your Property\033[0m")     // Green
		fmt.Println("\033[1;32m2. View Listed Property\033[0m")   // Green
		fmt.Println("\033[1;32m3. Chats\033[0m")                  // Green
		fmt.Println("\033[1;31m4. Back to Main Dashboard\033[0m") // Red
		fmt.Print("\nEnter your choice: ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Printf("\033[1;31mError reading input: %v\033[0m\n", err) // Red
			continue
		}

		switch choice {
		case 1:
			fmt.Println("\033[1;33m\nListing a new property...\033[0m") // Yellow
			// Add functionality to list a property here
			ui.ListPropertyUI()
		case 2:
			fmt.Println("\033[1;33m\nViewing listed properties...\033[0m") // Yellow
			// Add functionality to view listed properties here
			listedProperty, err := ui.propertyService.GetAllListedProperties()
			if err != nil {
				fmt.Printf("\033[1;31mError in fetching Listed Properties : %v\033[0m\n", err)
			}
			utils.DisplayProperties(listedProperty)

		case 3:
			fmt.Println("\033[1;33m\nOpening chats...\033[0m") // Yellow
			// Add functionality to handle chats here
		case 4:
			return // Go back to the main dashboard
		default:
			fmt.Println("\033[1;31m\nInvalid choice. Please select a valid option.\033[0m") // Red
		}
	}
}
