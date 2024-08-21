package ui

import (
	"fmt"
	"rentease/internal/domain/entities"
	"rentease/pkg/utils"
	"strings"
)

func (ui *UI) ListPropertyUI() {

	fmt.Print("Enter property type (1. Commercial, 2. House, 3. Flat): ")
	var propertyType int
	fmt.Scanf("%d", &propertyType)
	title := utils.ReadInput("\nEnter property title: ")

	// Collect common property details
	area := utils.ReadInput("Enter area: ")
	city := utils.ReadInput("Enter city: ")
	state := utils.ReadInput("Enter state: ")

	var pincode int
	pincode = utils.ReadPincode()
	landlordUsername := utils.ReadInput("Enter landlord username: ")

	var details interface{}
	switch propertyType {
	case 1:
		// Collect Commercial-specific details
		floorArea := utils.ReadInput("Enter floor area: ")
		subType := utils.ReadInput("Enter subtype (shop, factory, warehouse): ")

		details = entities.CommercialDetails{
			FloorArea: floorArea,
			SubType:   subType,
		}
	case 2:
		// Collect House-specific details
		fmt.Print("Enter number of rooms: ")
		var noOfRooms int
		fmt.Scanf("%d", &noOfRooms)

		furnishedCategory := utils.ReadInput("Enter furnished category: ")
		amenitiesStr := utils.ReadInput("Enter amenities (comma separated): ")
		amenities := strings.Split(amenitiesStr, ",")

		details = entities.HouseDetails{
			NoOfRooms:         noOfRooms,
			FurnishedCategory: furnishedCategory,
			Amenities:         amenities,
		}
	case 3:
		// Collect Flat-specific details
		furnishedCategory := utils.ReadInput("Enter furnished category: ")
		amenitiesStr := utils.ReadInput("Enter amenities (comma separated): ")
		amenities := strings.Split(amenitiesStr, ",")

		fmt.Print("Enter BHK: ")
		var bhk int
		fmt.Scanf("%d", &bhk)

		details = entities.FlatDetails{
			FurnishedCategory: furnishedCategory,
			Amenities:         amenities,
			BHK:               bhk,
		}
	default:
		fmt.Println("Invalid property type")
		return
	}

	property := entities.Property{
		PropertyType:     propertyType,
		Title:            title,
		Address:          entities.Address{Area: area, City: city, State: state, Pincode: pincode},
		LandlordUsername: landlordUsername,
		IsApproved:       false,
		Details:          details,
	}

	// Save property to the repository
	err := ui.propertyService.ListProperty(property)
	if err != nil {
		fmt.Println("Error listing property:", err)
	} else {
		fmt.Println("Property listed successfully.")
	}
}
