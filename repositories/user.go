package repositories

import "go.mongodb.org/mongo-driver/mongo"

type UserRepository struct {
	Database *mongo.Database
}

func NewUserRepository(database *mongo.Database) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}
