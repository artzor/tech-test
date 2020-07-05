// Package store provides APIs for data persistence
package store

import (
	"context"
	"log"
	"portdomain/entity"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Store provides save/retrieve
type Store struct {
	ports *mongo.Collection
}

// Connect establishes connection to mongodb instance
func Connect(dbInstance string, dbName string) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbInstance))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client.Database(dbName), nil
}

// New returns new Store instance
func New(db *mongo.Database) *Store {
	return &Store{
		ports: db.Collection("ports"),
	}
}

// Save creates document in mongodb
func (s *Store) Save(ctx context.Context, pd entity.PortDetails) error {
	log.Printf("[info] saving record: %+v", pd)
	_, err := s.ports.UpdateOne(ctx, bson.M{"_id": pd.ID}, bson.M{"$set": pd}, options.Update().SetUpsert(true))
	return err
}

// Get retrieves port details from mongodb by port id
func (s *Store) Get(ctx context.Context, portID string) (entity.PortDetails, error) {
	doc := s.ports.FindOne(ctx, bson.M{"_id": portID})
	if err := doc.Err(); err != nil {
		return entity.PortDetails{}, err
	}

	pd := entity.PortDetails{}
	if err := doc.Decode(&pd); err != nil {
		return entity.PortDetails{}, err
	}

	return pd, nil
}
