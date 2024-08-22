package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"rentease/internal/domain/entities"
)

type RequestRepoMongo struct {
	collection *mongo.Collection
}

func NewRequestRepoMongo(collection *mongo.Collection) *RequestRepoMongo {
	return &RequestRepoMongo{collection: collection}
}

func (r *RequestRepoMongo) SaveRequest(request entities.PropertyRequest) error {
	_, err := r.collection.InsertOne(context.Background(), request)
	return err
}

func (r *RequestRepoMongo) UpdateRequestStatus(requestID primitive.ObjectID, newStatus string) error {
	filter := bson.M{"_id": requestID}
	update := bson.M{"$set": bson.M{"requestStatus": newStatus}}

	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (r *RequestRepoMongo) FindRequestsByLandlord(landlordName string) ([]entities.PropertyRequest, error) {
	filter := bson.M{"landlordName": landlordName, "requestStatus": "Pending"}
	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var requests []entities.PropertyRequest
	for cursor.Next(context.Background()) {
		var request entities.PropertyRequest
		if err := cursor.Decode(&request); err != nil {
			return nil, err
		}
		requests = append(requests, request)
	}

	return requests, nil
}

func (r *RequestRepoMongo) FindRequestsByTenant(tenantName string) ([]entities.PropertyRequest, error) {
	filter := bson.M{"tenantName": tenantName}
	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var requests []entities.PropertyRequest
	for cursor.Next(context.Background()) {
		var request entities.PropertyRequest
		if err := cursor.Decode(&request); err != nil {
			return nil, err
		}
		requests = append(requests, request)
	}

	return requests, nil
}
