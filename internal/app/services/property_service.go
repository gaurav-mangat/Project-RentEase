package services

import (
	"rentease/internal/domain/entities"
	"rentease/internal/domain/interfaces"
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
