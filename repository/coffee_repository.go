package repository

import (
	db "github.com/sugigran17/golang-clean-architecture/database"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type coffeeRepository struct {
	Collection *mongo.Collection
}

type ICoffeeRepository interface {
}

func NewCoffeeRepository(db *db.MongoDB) ICoffeeRepository {
	return &coffeeRepository{
		Collection: db.MongoDB.Database(os.Getenv("DATABASE_NAME")).Collection("coffee"),
	}
}
