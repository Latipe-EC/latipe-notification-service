package mongodb

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"latipe-notification-service/config"
)

type MongoClient struct {
	client  *mongo.Client
	rootCtx context.Context
	cfg     *config.AppConfig
}

// Open - creates a new Mongo
func OpenMongoDBConnection(cfg *config.AppConfig) (*MongoClient, error) {
	ctx := context.Background()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.DB.Mongodb.Connection))
	if err != nil {
		return nil, err
	}

	return &MongoClient{client: client, rootCtx: ctx, cfg: cfg}, nil

}

func (m *MongoClient) GetDB() *mongo.Database {
	db := m.client.Database(m.cfg.DB.Mongodb.DbName)
	return db
}

// Disconnect - used mainly in testing to avoid capping out the concurrent connections on MongoDB
func (m *MongoClient) Disconnect() {
	err := m.client.Disconnect(m.rootCtx)
	if err != nil {
		log.Fatalf("disconnecting from mongodb: %v", err)
	}
}

// Ping sends a ping command to verify that the client can connect to the deployment.
func (m *MongoClient) Ping() error {
	return m.client.Ping(m.rootCtx, readpref.Primary())
}
