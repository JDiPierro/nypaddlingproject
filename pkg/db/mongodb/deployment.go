package mongodb

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/ricoberger/go-vue-starter/pkg/model"
	"github.com/sirupsen/logrus"
)

// CreateDeployment saves a Deployment model to the database
func (db *DB) CreateDeployment(deploy *model.Deployment) error {
	deploy.ID = primitive.NewObjectID().Hex()

	_, err := db.collections.deployments.InsertOne(context.Background(), deploy)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteErrors); ok {
			if len(writeErr) == 1 && writeErr[0].Code == 11000 {
				return fmt.Errorf("deploy_version_already_exists")
			}
		}
		return err
	}
	return nil
}

func (db *DB) GetDeployment(id string) (*model.Deployment, error) {
	return nil, nil
}

func (db *DB) GetDeployments() ([]*model.Deployment, error) {
	var deploys []*model.Deployment
	ctx := context.Background()
	cursor, err := db.collections.deployments.Find(ctx, bson.D{{}})
	if err != nil {
		logrus.WithError(err).Error("Deployment Find Error")
		return nil, err
	}
	if cursor.Err() != nil {
		logrus.WithError(err).Error("Deployments Cursor Error")
		return nil, cursor.Err()
	}

	for cursor.Next(ctx) {
		var deploy model.Deployment

		err := cursor.Decode(&deploy)
		if err != nil {
			logrus.WithError(err).Error("Deployment decode error")
		}
		deploys = append(deploys, &deploy)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return deploys, nil
}

func (db *DB) GetLatestDeployForAppEnv(appID string, envID string) (*model.Deployment, error) {
	var deploys []*model.Deployment
	ctx := context.Background()
	opts := &options.FindOptions{
		Sort: bson.M{
			"timestamp": -1,
		},
	}
	cursor, err := db.collections.deployments.Find(ctx, bson.M{"app_id": appID, "env_id": envID}, opts)
	if err != nil {
		logrus.WithError(err).Error("Deployment Find Error")
		return nil, err
	}
	if cursor.Err() != nil {
		logrus.WithError(err).Error("Deployments Cursor Error")
		return nil, cursor.Err()
	}

	for cursor.Next(ctx) {
		var deploy model.Deployment

		err := cursor.Decode(&deploy)
		if err != nil {
			logrus.WithError(err).Error("Deployment decode error")
		}
		deploys = append(deploys, &deploy)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(deploys) == 0 {
		return nil, nil
	}

	return deploys[0], nil
}

func (db *DB) GetDeploymentsForApp(appID string) ([]*model.Deployment, error) {
	return nil, nil
}

func (db *DB) GetDeploymentsForEnv(envID string) ([]*model.Deployment, error) {
	return nil, nil
}
