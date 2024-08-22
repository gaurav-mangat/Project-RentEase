package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"rentease/internal/domain/entities"
	"rentease/internal/domain/interfaces"
	"rentease/pkg/utils"
)

type PropertyRepo struct {
	client     *mongo.Client
	collection *mongo.Collection
}

// NewPropertyRepo initializes a new PropertyRepo with a MongoDB connection.
func NewPropertyRepo(uri string, dbName string, collectionName string) (interfaces.PropertyRepo, error) {
	client, err := connectToMongoDB(uri)
	if err != nil {
		return nil, err
	}

	collection := client.Database(dbName).Collection(collectionName)
	return &PropertyRepo{
		client:     client,
		collection: collection,
	}, nil
}

// SaveProperty saves a property to the MongoDB collection.
func (r *PropertyRepo) SaveProperty(property entities.Property) error {
	_, err := r.collection.InsertOne(context.TODO(), property)
	if err != nil {
		return err
	}
	return nil
}

// GetAllListedProperties retrieves all listed properties from the collection.
func (r *PropertyRepo) GetAllListedProperties() ([]entities.Property, error) {
	// Create a filter to match properties by the landlord's username
	filter := bson.D{{"landlordusername", utils.ActiveUser}}

	// Query the database with the filter
	cursor, err := r.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to query properties: %w", err)
	}
	defer cursor.Close(context.TODO())

	var properties []entities.Property
	for cursor.Next(context.TODO()) {
		var property entities.Property
		// Decode the base property structure first
		if err := cursor.Decode(&property); err != nil {
			return nil, fmt.Errorf("failed to decode property: %w", err)
		}

		// Decode the Details field based on PropertyType
		switch property.PropertyType {
		case 1: // Commercial
			var details entities.CommercialDetails
			if err := bson.Unmarshal(cursor.Current.Lookup("details").Value, &details); err != nil {
				return nil, fmt.Errorf("failed to decode commercial details: %w", err)
			}
			property.Details = details
		case 2: // House
			var details entities.HouseDetails
			if err := bson.Unmarshal(cursor.Current.Lookup("details").Value, &details); err != nil {
				return nil, fmt.Errorf("failed to decode house details: %w", err)
			}
			property.Details = details
		case 3: // Flat
			var details entities.FlatDetails
			if err := bson.Unmarshal(cursor.Current.Lookup("details").Value, &details); err != nil {
				return nil, fmt.Errorf("failed to decode flat details: %w", err)
			}
			property.Details = details
		default:
			// Unknown property type, Details will remain nil
		}

		properties = append(properties, property)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return properties, nil
}

// UpdateListedProperty updates an existing property in the collection.
func (r *PropertyRepo) UpdateListedProperty(property entities.Property) error {
	filter := bson.D{{"title", property.Title}}
	update := bson.D{
		{"$set", bson.D{
			{"address", property.Address},
			{"landlordusername", property.LandlordUsername},
			{"isapproved", property.IsApproved},
			{"details", property.Details},
		}},
	}

	// If the property is approved, reset its approval status to false
	if property.IsApproved {
		property.IsApproved = false
	}

	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

// DeleteListedProperty deletes a property from the collection by title.
func (r *PropertyRepo) DeleteListedProperty(propertyTitle string) error {
	filter := bson.D{{"title", propertyTitle}}
	_, err := r.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
