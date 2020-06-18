package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/ricoberger/go-vue-starter/pkg/model"
	"github.com/sirupsen/logrus"
)

// CreateApplication saves a new Application model to the database
func (db *DB) CreateApplication(app *model.Application) error {
	return db.SaveApplication(app)
}

// SaveApplication saves an existing Application model to the database
func (db *DB) SaveApplication(app *model.Application) error {
	app.ID = app.Name

	item, err := dynamodbattribute.MarshalMap(app)
	if err != nil {
		return err
	}

	_, err = db.client.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(db.ApplicationsTableName()),
	})
	if err != nil {
		logrus.Debug("SaveApplication error")
	}

	return err
}

func (db *DB) GetApplications() ([]*model.Application, error) {
	// Query Dynamo for events since the timestamp provided
	output, err := db.client.Scan(&dynamodb.ScanInput{
		ConsistentRead: aws.Bool(true),
		TableName:      aws.String(db.ApplicationsTableName()),
	})
	if err != nil {
		return nil, err
	} else if output == nil || *output.Count == 0 {
		return nil, nil
	}

	// Unmarshal the messages pulled from Dynamo
	results := make([]*model.Application, 0)
	for _, item := range output.Items {
		app := &model.Application{}
		err := dynamodbattribute.UnmarshalMap(item, app)
		if err != nil {
			logrus.WithError(err).Warn("Application Unmarshalling Error")
			continue
		}
		results = append(results, app)
	}

	return results, err
}

func (db *DB) BulkGetApplications(appIDs []string) ([]*model.Application, error) {
	// Build list of App IDs to grab
	reqKeys := make([]map[string]*dynamodb.AttributeValue, 0)
	for _, appID := range appIDs {
		reqKeys = append(reqKeys, map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(appID),
			},
		})
	}

	// Query Dynamo
	output, err := db.client.BatchGetItem(&dynamodb.BatchGetItemInput{
		RequestItems: map[string]*dynamodb.KeysAndAttributes{
			db.ApplicationsTableName(): {
				Keys: reqKeys,
			},
		},
	})
	if err != nil {
		return nil, err
	} else if output == nil || output.Responses == nil {
		logrus.Warn("No applications loaded from batch get")
		return nil, nil
	}

	// Unmarshal the messages pulled from Dynamo
	results := make([]*model.Application, 0)
	for _, item := range output.Responses[db.ApplicationsTableName()] {
		app := &model.Application{}
		err := dynamodbattribute.UnmarshalMap(item, app)
		if err != nil {
			logrus.WithError(err).Warn("Application Unmarshalling Error")
			continue
		}
		results = append(results, app)
	}

	return results, err
}

func (db *DB) GetApplication(id string) (*model.Application, error) {
	return nil, nil
}

func (db *DB) GetApplicationByName(name string) (*model.Application, error) {
	output, err := db.client.GetItem(&dynamodb.GetItemInput{
		ConsistentRead: aws.Bool(true),
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(name),
			},
		},
		TableName: aws.String(db.ApplicationsTableName()),
	})

	if err != nil {
		return nil, err
	} else if output == nil || output.Item == nil {
		return nil, nil
	}

	result := &model.Application{}
	err = dynamodbattribute.UnmarshalMap(output.Item, result)
	return result, err
}
