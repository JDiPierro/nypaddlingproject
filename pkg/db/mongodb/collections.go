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
	apps := db.Collection("applications")
	_, err = apps.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.M{
			"name": 1,
		},
		Options: uniqueIndexOptions,
	})
	if err != nil {
		logrus.WithError(err).Error("Unable to create apps index")
	}

	// Envs
	envs := db.Collection("environments")
	_, err = envs.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.M{
			"name": 1,
		},
		Options: uniqueIndexOptions,
	})
	if err != nil {
		logrus.WithError(err).Error("Unable to create envs index")
	}

	// Deploys
	deploys := db.Collection("deployments")
	_, err = deploys.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.M{
			"app_id": 1,
			"env_id": 1,
			"version": 1,
		},
		Options: uniqueIndexOptions,
	})
	if err != nil {
		logrus.WithError(err).Error("Unable to create deploys index")
	}

	collections := &Collections{
		users: users,
		applications: apps,
		environments: envs,
		deployments: deploys,
	}

	return collections
}
