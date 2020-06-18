package mongodb

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Config represents MongoDB configuration
type Config struct {
	ConnectionURI string `yaml:"connection_uri"`
	DatabaseName  string `yaml:"database_name"`
}

// DB represents the structure of the database
type DB struct {
	config      *Config
	client      *mongo.Client
	collections *Collections
}

// Collections represents all needed db collections
type Collections struct {
	users *mongo.Collection
	applications *mongo.Collection
	environments *mongo.Collection
	deployments *mongo.Collection
}

// NewConnection creates a new database connection
func NewConnection(config *Config) (*DB, error) {
	client, err := mongo.Connect(context.Background(), config.ConnectionURI)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	collections := initCollections(client, config)

	return &DB{
		config:      config,
		client:      client,
		collections: collections,
	}, nil
}

// CloseConnection closes the database connection
func (db *DB) CloseConnection() error {
	err := db.client.Disconnect(context.Background())
	if err != nil {
		return err
	}

	return nil
}
