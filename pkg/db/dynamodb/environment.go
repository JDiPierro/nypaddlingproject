package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/ricoberger/go-vue-starter/pkg/model"
	"github.com/sirupsen/logrus"
)

// CreateEnvironment saves an Environment model to the database
func (db *DB) CreateEnvironment(env *model.Environment) error {
	env.ID = env.Name

	item, err := dynamodbattribute.MarshalMap(env)
	if err != nil {
		return err
	}

	_, err = db.client.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(db.EnvironmentsTableName()),
	})
	if err != nil {
		logrus.Debug("CreateEnvironment error")
	}

	return err
}

func (db *DB) GetEnvironments() ([]*model.Environment, error) {
	// Query Dynamo for events since the timestamp provided
	output, err := db.client.Scan(&dynamodb.ScanInput{
		ConsistentRead: aws.Bool(true),
		TableName:      aws.String(db.EnvironmentsTableName()),
	})
	if err != nil {
		return nil, err
	} else if output == nil || *output.Count == 0 {
		return nil, nil
	}

	// Unmarshal the messages pulled from Dynamo
	envs := make([]*model.Environment, 0)
	for _, item := range output.Items {
		env := &model.Environment{}
		err := dynamodbattribute.UnmarshalMap(item, env)
		if err != nil {
			logrus.WithError(err).Warn("Error unmarshaling Deployment dynamo doc")
			continue
		}
		envs = append(envs, env)
	}

	return envs, err
}

func (db *DB) GetEnvironmentsForApp(appID string) ([]*model.Environment, error) {
	return nil, nil
}

func (db *DB) GetEnvironmentByName(name string) (*model.Environment, error) {
	output, err := db.client.GetItem(&dynamodb.GetItemInput{
		ConsistentRead: aws.Bool(true),
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(name),
			},
		},
		TableName: aws.String(db.EnvironmentsTableName()),
	})

	if err != nil {
		return nil, err
	} else if output == nil || output.Item == nil {
		return nil, nil
	}

	result := &model.Environment{}
	err = dynamodbattribute.UnmarshalMap(output.Item, result)
	return result, err
}
