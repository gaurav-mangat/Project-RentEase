package ui

import (
	"fmt"
	"rentease/pkg/utils"
	"strings"
)

func (ui *UI) SearchPropertyUI() {
	fmt.Println("\n\033[1;34m========================\033[0m") // Blue
	fmt.Println("\033[1;34mSearch Property\033[0m")            // Blue
	fmt.Println("\033[1;34m========================\033[0m")   // Blue

	// Collect search criteria from the tenant
	var propertyType int
	for {
		fmt.Print("Enter property type (1. Commercial, 2. House, 3. Flat): ")
		_, err := fmt.Scanf("%d", &propertyType)

		// Check if the input is valid
		if err != nil || propertyType < 1 || propertyType > 3 {
			fmt.Println("\033[1;31mInvalid input. Please enter a valid property type (1, 2, or 3).\033[0m") // Red
			continue
		}

		// Break the loop if the input is valid
		break
	}

	area := utils.ReadInput("Enter locality (leave blank to skip): ")
	city := utils.ReadInput("Enter city (leave blank to skip): ")
	state := utils.ReadInput("Enter state (leave blank to skip): ")
	pincode := utils.ReadPincode()

	// Call service to search for properties
	properties, err := ui.propertyService.SearchProperties(area, city, state, pincode, propertyType)
	if err != nil {
		fmt.Printf("\033[1;31mError searching properties: %v\033[0m\n", err) // Red
		return
	}

	// Display the search results
	if len(properties) == 0 {
		fmt.Println("\033[1;33mNo properties found matching your criteria.\033[0m") // Yellow
		return
	}

	fmt.Println("\n\033[1;34mSearch Results\033[0m")         // Blue
	fmt.Println("\033[1;34m========================\033[0m") // Blue
	utils.DisplayPropertyshortInfo(properties)

	for {
		fmt.Print("Enter the property number to see more details (or 0 to exit): ")
		var choice int
		fmt.Scan(&choice)

		if choice == 0 {
			break
		}

		if choice < 1 || choice > len(properties) {
			fmt.Println("\033[1;31mInvalid property number.\033[0m") // Red
			continue
		}

		prop := properties[choice-1]

		utils.DisplayProperty(prop)

		// Fetch landlord details
		landlord, err := ui.userService.FindByUsername(prop.LandlordUsername)
		if err != nil {
			fmt.Printf("\033[1;31mError fetching landlord details: %v\033[0m\n", err) // Red
			continue
		}

		fmt.Println("\nLandlord Details are:")
		fmt.Printf("  Landlord Name: %s\n", landlord.Name)
		fmt.Printf("  Landlord Phone: %s\n", landlord.PhoneNumber)
		fmt.Printf("  Landlord Email: %s\n", landlord.Email)
		fmt.Printf("  Landlord Address: %s\n", landlord.Address)

		// Option to add property to wishlist
		fmt.Print("\nWould you like to add this property to your wishlist? (yes/no): ")
		var addToWishlist string
		fmt.Scan(&addToWishlist)
		addToWishlist = strings.ToLower(addToWishlist)

		if addToWishlist == "yes" {
			err := ui.userService.AddToWishlist(utils.ActiveUser, prop.ID)
			if err != nil {
				fmt.Printf("\033[1;31mError adding property to wishlist: %v\033[0m\n", err) // Red
			} else {
				fmt.Println("\033[1;32mProperty added to wishlist successfully.\033[0m") // Green
			}
		}
		// Option to request property
		fmt.Print("\nWould you like to request this property? (yes/no): ")
		var request string
		fmt.Scan(&request)
		request = strings.ToLower(request)

		if request == "yes" {
			if utils.ActiveUser != prop.LandlordUsername {
				err = ui.requestService.CreatePropertyRequest(utils.ActiveUser, prop.ID, prop.LandlordUsername)
				if err != nil {
					fmt.Printf("\033[1;31mError requesting property: %v\033[0m\n", err) // Red
				} else {
					fmt.Println("\033[1;32mProperty request sent successfully.\033[0m") //  Green
				}
			} else {
				fmt.Println("You can't make a request for your property !")
			}
		}

		// Option to see details of another property
		fmt.Print("\nWould you like to see details of another property? (yes/no): ")
		var response string
		fmt.Scan(&response)
		response = strings.ToLower(response)

		if response != "yes" {
			break
		}
	}
}
