package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/ricoberger/go-vue-starter/pkg/model"
	"github.com/sirupsen/logrus"
)

type DeploymentDynamoDocument struct {
	Model         *model.Deployment `json:"deploy"`
	AppEnvVersion string            `json:"app_env_version"`
	AppEnv        string            `json:"app_env"`
	EnvID         string            `json:"env_id"`
	Timestamp     string            `json:"timestamp"`
}

func dynamoDocForDeployment(deploy *model.Deployment) *DeploymentDynamoDocument {
	deploy.ID = deploy.Version

	return &DeploymentDynamoDocument{
		Model:         deploy,
		AppEnvVersion: deploy.AppID + "::" + deploy.EnvID + "::" + deploy.Version,
		AppEnv:        deploy.AppID + "::" + deploy.EnvID,
		EnvID:         deploy.EnvID,
		Timestamp:     deploy.Timestamp,
	}
}

// CreateDeployment saves a Deployment model to the database
func (db *DB) CreateDeployment(deploy *model.Deployment) error {
	doc := dynamoDocForDeployment(deploy)
	item, err := dynamodbattribute.MarshalMap(doc)
	if err != nil {
		return err
	}

	_, err = db.client.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(db.DeploymentsTableName()),
		/* Don't enforce uniqueness so redeploys update the timestamp of that app_env_version
		ConditionExpression: aws.String("attribute_not_exists(app_env_version)"),
		*/
	})
	if err != nil {
		logrus.Debug("CreateDeployment error")
	}

	return err
}

func (db *DB) GetDeployment(id string) (*model.Deployment, error) {
	return nil, nil
}

func (db *DB) GetDeployments() ([]*model.Deployment, error) {
	// Query Dynamo for events since the timestamp provided
	output, err := db.client.Scan(&dynamodb.ScanInput{
		ConsistentRead: aws.Bool(true),
		TableName:      aws.String(db.DeploymentsTableName()),
	})
	if err != nil {
		return nil, err
	} else if output == nil || *output.Count == 0 {
		return nil, nil
	}

	deployments := make([]*model.Deployment, 0)
	for _, item := range output.Items {
		doc := &DeploymentDynamoDocument{}
		err := dynamodbattribute.UnmarshalMap(item, doc)
		if err != nil {
			logrus.WithError(err).Warn("Error unmarshaling Deployment dynamo doc")
			continue
		}
		if doc.Model == nil {
			logrus.Error("Model object missing from unmarshalled dynamo document")
		}
		deployments = append(deployments, doc.Model)
	}

	return deployments, err
}

func (db *DB) GetLatestDeployForAppEnv(appID string, envID string) (*model.Deployment, error) {
	// Query Dynamo for events since the timestamp provided
	output, err := db.client.Query(&dynamodb.QueryInput{
		TableName:              aws.String(db.DeploymentsTableName()),
		IndexName:              aws.String("app_env_timestamp"),
		KeyConditionExpression: aws.String("app_env = :app_env"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":app_env": {
				S: aws.String(appID + "::" + envID),
			},
		},
		Limit: aws.Int64(1),
	})
	if err != nil {
		return nil, err
	} else if output == nil || *output.Count == 0 {
		return nil, nil
	}

	// Unmarshal the messages pulled from Dynamo
	deployments := make([]*model.Deployment, 0)
	for _, item := range output.Items {
		doc := &DeploymentDynamoDocument{}
		err := dynamodbattribute.UnmarshalMap(item, doc)
		if err != nil {
			logrus.WithError(err).Warn("Error unmarshaling Deployment dynamo doc")
			continue
		}
		if doc.Model == nil {
			logrus.Error("Model object missing from unmarshalled dynamo document")
		}
		deployments = append(deployments, doc.Model)
	}

	return deployments[0], err
}

func (db *DB) GetDeploymentsForApp(appID string) ([]*model.Deployment, error) {
	return nil, nil
}

func (db *DB) GetDeploymentsForEnv(envID string) ([]*model.Deployment, error) {
	// Query Dynamo for events since the timestamp provided
	output, err := db.client.Query(&dynamodb.QueryInput{
		TableName:              aws.String(db.DeploymentsTableName()),
		IndexName:              aws.String("env_deploys"),
		KeyConditionExpression: aws.String("env_id = :env_id"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":env_id": {
				S: aws.String(envID),
			},
		},
	})
	if err != nil {
		return nil, err
	} else if output == nil || *output.Count == 0 {
		return nil, nil
	}

	// Unmarshal the messages pulled from Dynamo
	deployments := make([]*model.Deployment, 0)
	for _, item := range output.Items {
		doc := &DeploymentDynamoDocument{}
		err := dynamodbattribute.UnmarshalMap(item, doc)
		if err != nil {
			logrus.WithError(err).Warn("Error unmarshaling Deployment dynamo doc")
			continue
		}
		if doc.Model == nil {
			logrus.Error("Model object missing from unmarshalled dynamo document")
		}
		deployments = append(deployments, doc.Model)
	}

	return deployments, err
}
