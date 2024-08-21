package entities

// Abstract Class for User
type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	Address      string `json:"address"`
	Role         string `json:"role"`
}

type NormalUser struct {
	User
	Wishlist        []int `json:"wishlist"`
	PropertyListing []int `json:"propertyListing"`
}

type AdminUser struct {
	User
}
