package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rentease/internal/domain/entities"
	"rentease/internal/domain/interfaces"
)

type RequestService struct {
	requestRepo interfaces.RequestRepo
}

func NewRequestService(requestRepo interfaces.RequestRepo) *RequestService {
	return &RequestService{
		requestRepo: requestRepo,
	}
}

func (rs *RequestService) CreatePropertyRequest(tenantName string, propertyID primitive.ObjectID, landlordName string) error {
	request := entities.PropertyRequest{
		TenantName:    tenantName,
		PropertyID:    propertyID,
		LandlordName:  landlordName,
		RequestStatus: "Pending",
	}

	return rs.requestRepo.SaveRequest(request)
}

func (rs *RequestService) UpdateRequestStatus(requestID primitive.ObjectID, newStatus string) error {
	return rs.requestRepo.UpdateRequestStatus(requestID, newStatus)
}

func (rs *RequestService) GetRequestsForLandlord(landlordName string) ([]entities.PropertyRequest, error) {
	return rs.requestRepo.FindRequestsByLandlord(landlordName)
}

func (rs *RequestService) GetRequestsForTenant(tenantName string) ([]entities.PropertyRequest, error) {
	return rs.requestRepo.FindRequestsByTenant(tenantName)
}
