package mongodb

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/sirupsen/logrus"
)

func initCollections(client *mongo.Client, config *Config) *Collections {
	db := client.Database(config.DatabaseName)

	uniqueIndexOptions := options.Index()
	uniqueIndexOptions.SetUnique(true)

	// Users
	users := db.Collection("users")
	_, err := users.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.M{
			"email": 1,
		},
		Options: uniqueIndexOptions,
	})
	if err != nil {
		logrus.WithError(err).Error("Unable to create users index")
	}

	// Apps
	locs := db.Collection("locations")
	_, err = locs.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.M{
			"name": 1,
		},
		Options: uniqueIndexOptions,
	})
	if err != nil {
		logrus.WithError(err).Error("Unable to create locations index")
	}



	collections := &Collections{
		users: users,
		locations: locs,
	}

	return collections
}
