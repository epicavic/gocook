package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Item represents an item at a shop
type Item struct {
	Name  string
	Price int64
}

// Storage is our storage interface
type Storage interface {
	GetByName(context.Context, string) (*Item, error)
	Put(context.Context, *Item) error
}

// MongoStorage implements our storage interface
type MongoStorage struct {
	*mongo.Client        // db connection
	DB            string // db name
	Collection    string // db collection (table)
}

// Put adds an item to our mongo instance
func (m *MongoStorage) Put(ctx context.Context, i *Item) error {
	c := m.Client.Database(m.DB).Collection(m.Collection)
	_, err := c.InsertOne(ctx, i)
	return err
}

// GetByName queries mongodb for an item with the correct name
func (m *MongoStorage) GetByName(ctx context.Context, name string) (*Item, error) {
	c := m.Client.Database(m.DB).Collection(m.Collection)
	var i Item
	if err := c.FindOne(ctx, bson.M{"name": name}).Decode(&i); err != nil {
		return nil, err
	}

	return &i, nil
}
