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

// CreateApplication saves a new Application model to the database
func (db *DB) CreateApplication(app *model.Application) error {
	app.ID = primitive.NewObjectID().Hex()

	_, err := db.collections.applications.InsertOne(context.Background(), app)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteErrors); ok {
			if len(writeErr) == 1 && writeErr[0].Code == 11000 {
				return fmt.Errorf("app_name_already_exists")
			}
		}
		return err
	}
	return nil
}

// SaveApplication saves an existing Application model to the database
func (db *DB) SaveApplication(app *model.Application) error {
	cursor := db.collections.applications.FindOneAndReplace(
		context.Background(),
		bson.D{primitive.E{
			Key:   "_id",
			Value: app.ID,
		}},
		app,
	)

	if cursor.Err() != nil {
		if writeErr, ok := cursor.Err().(mongo.WriteErrors); ok {
			if len(writeErr) == 1 && writeErr[0].Code == 11000 {
				return fmt.Errorf("app_name_already_exists")
			}
		}

		return cursor.Err()
	}

	return nil
}

func (db *DB) GetApplications() ([]*model.Application, error) {
	var apps []*model.Application
	ctx := context.Background()
	cursor, err := db.collections.applications.Find(ctx, bson.D{{}})
	if err != nil {
		logrus.WithError(err).Error("Application Find Error")
		return nil, err
	}
	if cursor.Err() != nil {
		logrus.WithError(err).Error("Applications Cursor Error")
		return nil, cursor.Err()
	}

	for cursor.Next(ctx) {
		var app model.Application

		err := cursor.Decode(&app)
		if err != nil {
			logrus.WithError(err).Error("Application decode error")
		}
		apps = append(apps, &app)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return apps, nil
}

func (db *DB) BulkGetApplications(appIDs []string) ([]*model.Application, error) {
	return nil, nil
}

func (db *DB) GetApplication(id string) (*model.Application, error) {
	return nil, nil
}

func (db *DB) GetApplicationByName(name string) (*model.Application, error) {
	app := &model.Application{}

	cursor := db.collections.applications.FindOne(
		context.Background(),
		bson.D{primitive.E{
			Key:   "name",
			Value: name,
		}},
	)

	if cursor.Err() != nil {
		return nil, cursor.Err()
	}

	err := cursor.Decode(app)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return app, nil
}
