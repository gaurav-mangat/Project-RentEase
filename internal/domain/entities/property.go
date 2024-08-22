package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Property struct {
	ID               primitive.ObjectID `json:"_id,omitempty"` // MongoDB unique ID
	PropertyType     int                `json:"property_type"` // 1: Commercial, 2: House, 3: Flat
	Title            string             `json:"title"`
	Address          Address            `json:"address"`
	LandlordUsername string             `json:"landlord_username"`
	RentAmount       float64            `json:"rent_amount"`
	IsApproved       bool               `json:"is_approved"`
	Details          interface{}        `json:"details"` // Holds specific details based on property type
}

type Address struct {
	Area    string `json:"area"`
	City    string `json:"city"`
	State   string `json:"state"`
	Pincode int    `json:"pincode"`
}

type CommercialDetails struct {
	FloorArea string `json:"floor_area"`
	SubType   string `json:"sub_type"` // shop, factory, warehouse
}

type HouseDetails struct {
	NoOfRooms         int      `json:"no_of_rooms"`
	FurnishedCategory string   `json:"furnished_category"`
	Amenities         []string `json:"amenities"`
}

type FlatDetails struct {
	FurnishedCategory string   `json:"furnished_category"`
	Amenities         []string `json:"amenities"`
	BHK               int      `json:"bhk"`
}
