package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rentease/internal/domain/entities"
)

type RequestRepo interface {
	SaveRequest(request entities.PropertyRequest) error
	UpdateRequestStatus(requestID primitive.ObjectID, newStatus string) error
	FindRequestsByLandlord(landlordName string) ([]entities.PropertyRequest, error)
	FindRequestsByTenant(tenantName string) ([]entities.PropertyRequest, error)
}
