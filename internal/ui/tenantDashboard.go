package ui

import (
	"fmt"
	"rentease/pkg/utils"
)

func (ui *UI) TenantDashboardUI() {
	for {
		fmt.Println("\033[1;34m\n========================\033[0m") // Blue
		fmt.Println("\033[1;34mTenant Dashboard\033[0m")           // Blue
		fmt.Println("\033[1;34m========================\033[0m")   // Blue
		fmt.Println("1. Search Property")
		fmt.Println("2. Your Wishlist")
		fmt.Println("3. Logout")

		choice := utils.ReadInput("\nEnter your choice: ")

		switch choice {
		case "1":
			ui.SearchPropertyUI()
		case "2":
			//ui.ViewWishlistUI()
		case "3":
			fmt.Println("\033[1;32mLogging out...\033[0m") // Green
			return
		default:
			fmt.Println("\033[1;31mInvalid choice, please try again.\033[0m") // Red
		}
	}
}
