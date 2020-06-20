package mongodb

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/ricoberger/go-vue-starter/pkg/model"
	"github.com/sirupsen/logrus"
)

// CreateEnvironment saves an Environment model to the database
func (db *DB) CreateLocation(loc *model.Location) error {
	loc.ID = primitive.NewObjectID().Hex()

	_, err := db.collections.locations.InsertOne(context.Background(), loc)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteErrors); ok {
			if len(writeErr) == 1 && writeErr[0].Code == 11000 {
				return fmt.Errorf("env_name_already_exists")
			}
		}
		return err
	}
	return nil
}

func (db *DB) GetLocations() ([]*model.Location, error) {
	var envs []*model.Location
	ctx := context.Background()
	cursor, err := db.collections.locations.Find(ctx, bson.D{{}})
	if err != nil {
		logrus.WithError(err).Error("Environment Find Error")
		return nil, err
	}
	if cursor.Err() != nil {
		logrus.WithError(err).Error("Environment Cursor Error")
		return nil, cursor.Err()
	}

	for cursor.Next(ctx) {
		var loc model.Location

		err := cursor.Decode(&loc)
		if err != nil {
			logrus.WithError(err).Error("Environment decode error")
		}
		envs = append(envs, &loc)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return envs, nil
}

