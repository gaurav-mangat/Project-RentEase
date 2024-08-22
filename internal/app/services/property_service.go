package services

import (
	"rentease/internal/domain/entities"
	"rentease/internal/domain/interfaces"
	"strings"
)

type PropertyService struct {
	propertyRepo interfaces.PropertyRepo
}

func NewPropertyService(propertyRepo interfaces.PropertyRepo) *PropertyService {
	return &PropertyService{
		propertyRepo: propertyRepo,
	}
}

// ListProperty saves a property to the repository.
func (ps *PropertyService) ListProperty(property entities.Property) error {
	return ps.propertyRepo.SaveProperty(property)
}

// GetAllListedProperties retrieves all listed properties from the repository.
func (ps *PropertyService) GetAllListedProperties() ([]entities.Property, error) {
	return ps.propertyRepo.GetAllListedProperties()
}

// UpdateListedProperty updates a property in the repository.
func (ps *PropertyService) UpdateListedProperty(property entities.Property) error {
	// Check if the property is approved before updating
	if property.IsApproved {
		// Reset approval status if the property was approved
		property.IsApproved = false
	}
	return ps.propertyRepo.UpdateListedProperty(property)
}

// DeleteListedProperty deletes a property from the repository by ID.
func (ps *PropertyService) DeleteListedProperty(propertyID string) error {
	return ps.propertyRepo.DeleteListedProperty(propertyID)
}

// SearchProperties searches for properties based on the given criteria.
func (ps *PropertyService) SearchProperties(area, city, state string, pincode, propertyType int) ([]entities.Property, error) {
	properties, err := ps.propertyRepo.GetAllListedProperties()
	if err != nil {
		return nil, err
	}

	// Normalize the input strings
	area = strings.TrimSpace(strings.ToLower(area))
	city = strings.TrimSpace(strings.ToLower(city))
	state = strings.TrimSpace(strings.ToLower(state))

	var results []entities.Property
	for _, property := range properties {
		if property.PropertyType == propertyType {
			// Normalize the property address fields
			propArea := strings.TrimSpace(strings.ToLower(property.Address.Area))
			propCity := strings.TrimSpace(strings.ToLower(property.Address.City))
			propState := strings.TrimSpace(strings.ToLower(property.Address.State))

			if (strings.Contains(propArea, area) && strings.Contains(propCity, city) && property.Address.Pincode == pincode) ||
				(strings.Contains(propCity, city) && property.Address.Pincode == pincode) ||
				(property.Address.Pincode == pincode) ||
				(strings.Contains(propState, state)) {
				results = append(results, property)
			}
		}
	}

	return results, nil
}
