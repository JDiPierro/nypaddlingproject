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
func (db *DB) CreateEnvironment(env *model.Environment) error {
	env.ID = primitive.NewObjectID().Hex()

	_, err := db.collections.environments.InsertOne(context.Background(), env)
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

func (db *DB) GetEnvironments() ([]*model.Environment, error) {
	var envs []*model.Environment
	ctx := context.Background()
	cursor, err := db.collections.environments.Find(ctx, bson.D{{}})
	if err != nil {
		logrus.WithError(err).Error("Environment Find Error")
		return nil, err
	}
	if cursor.Err() != nil {
		logrus.WithError(err).Error("Environment Cursor Error")
		return nil, cursor.Err()
	}

	for cursor.Next(ctx) {
		var env model.Environment

		err := cursor.Decode(&env)
		if err != nil {
			logrus.WithError(err).Error("Environment decode error")
		}
		envs = append(envs, &env)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return envs, nil
}

func (db *DB) GetEnvironmentsForApp(appID string) ([]*model.Environment, error) {
	return nil, nil
}

func (db *DB) GetEnvironmentByName(name string) (*model.Environment, error) {
	env := &model.Environment{}

	cursor := db.collections.environments.FindOne(
		context.Background(),
		bson.D{primitive.E{
			Key:   "name",
			Value: name,
		}},
	)

	if cursor.Err() != nil {
		return nil, cursor.Err()
	}

	err := cursor.Decode(env)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return env, nil
}
