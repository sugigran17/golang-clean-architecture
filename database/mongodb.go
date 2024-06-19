package database

import (
	"context"
	"go.elastic.co/apm/module/apmmongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

type MongoDB struct {
	Context context.Context
	MongoDB *mongo.Client
}

type IMongoDB interface {
	Connect(context.Context) error
	Disconnect(context.Context) error
	Ping(context.Context) error
}

func NewMongoClient(connectionURL string, maxPoolSize uint64) *MongoDB {
	option := options.Client().ApplyURI(connectionURL).SetMonitor(apmmongo.CommandMonitor()).SetMaxPoolSize(maxPoolSize)

	client, err0 := mongo.Connect(context.Background(), option)

	if err0 != nil {
		log.Fatal(err0)
	}

	return &MongoDB{
		Context: context.Background(),
		MongoDB: client,
	}

}

// Disconnect Disconnect
func (ds MongoDB) Disconnect() {
	_ = ds.MongoDB.Disconnect(ds.Context)
}

// Ping Ping
func (ds MongoDB) Ping() {
	_ = ds.MongoDB.Ping(ds.Context, readpref.Primary())
}
