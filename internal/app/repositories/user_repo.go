package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"rentease/internal/domain/entities"
	"rentease/internal/domain/interfaces"
	"rentease/pkg/utils"
)

type UserRepo struct {
	client     *mongo.Client
	collection *mongo.Collection
}

// NewUserRepo initializes a new UserRepo with a MongoDB connection.
func NewUserRepo(uri string, dbName string, collectionName string) (interfaces.UserRepo, error) {
	client, err := connectToMongoDB(uri)
	if err != nil {
		return nil, err
	}

	collection := client.Database(dbName).Collection(collectionName)
	return &UserRepo{
		client:     client,
		collection: collection,
	}, nil
}

// connectToMongoDB creates a new MongoDB client and connects to the database.
func connectToMongoDB(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to MongoDB!")
	return client, nil
}

// SaveUser saves a user to the MongoDB collection.
func (r *UserRepo) SaveUser(user entities.User) error {
	_, err := r.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) FindByUsername(ctx context.Context, username string) (*entities.User, error) {
	var user entities.User
	filter := bson.D{{"username", username}}

	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No document found
		}
		return nil, fmt.Errorf("failed to find user by username: %w", err) // Wrap other errors
	}

	return &user, nil
}

// CheckPassword verifies the user's password.
func (r *UserRepo) CheckPassword(ctx context.Context, username, password string) (bool, error) {
	user, err := r.FindByUsername(ctx, username)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, nil // User not found
	}

	// Compare the provided password with the stored hashed password
	return utils.CheckPasswordHash(password, user.PasswordHash), nil
}

//func (r *UserRepo) FindByEmail(ctx context.Context, email string) (*entities.User, error) {}
