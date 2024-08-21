package interfaces

import (
	"rentease/internal/domain/entities"
)

type PropertyRepo interface {
	SaveProperty(property entities.Property) error
	GetAllListedProperties() ([]entities.Property, error)
	UpdateListedProperty(property entities.Property) error
	DeleteListedProperty(propertyID string) error
}

//type PropertyService interface {
//	ListProperty(property entities.Property) error
//	GetAllListedProperties(username string) ([]entities.Property, error)
//}
