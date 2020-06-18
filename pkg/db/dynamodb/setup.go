package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/sirupsen/logrus"
)

var hasCreatedUsers = false
var hasCreatedApplications = false
var hasCreatedEnvironments = false
var hasCreatedDeployments = false

func (db *DB) UsersTableName() string {
	return db.tableName("users")
}

func (db *DB) ApplicationsTableName() string {
	return db.tableName("applications")
}

func (db *DB) EnvironmentsTableName() string {
	return db.tableName("environments")
}

func (db *DB) DeploymentsTableName() string {
	return db.tableName("deployments")
}

func (db *DB) tableName(model string) string {
	return db.cfg.Env + "-biome-" + model
}

// SeedTables creates all the tables for the local environment
func (db *DB) SeedTables() error {

	for _, seedFn := range []func() error{
		db.createUsersTable,
		db.createApplicationsTable,
		db.createEnvironmentsTable,
		db.createDeploymentsTable,
	} {
		if err := seedFn(); err != nil {
			return err
		}
	}
	return nil
}

// CreateDynamoTableIfNotExists will create a DynamoDB table if it does not already exist
func CreateDynamoTableIfNotExists(client *dynamodb.DynamoDB, schema *dynamodb.CreateTableInput) error {
	return createTable(client, schema, true)
}

func createTable(client *dynamodb.DynamoDB, schema *dynamodb.CreateTableInput, ignoreExistingTableError bool) error {
	_, err := client.CreateTable(schema)

	if err != nil {
		aerr, ok := err.(awserr.Error)
		if ok && ignoreExistingTableError && aerr.Code() == dynamodb.ErrCodeResourceInUseException {
			err = nil
		}
	}

	return err
}

func (db *DB) createUsersTable() error {
	if hasCreatedUsers {
		return nil
	}

	err := CreateDynamoTableIfNotExists(db.client, &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{{
			AttributeName: aws.String("email"),
			AttributeType: aws.String(dynamodb.ScalarAttributeTypeS),
		}},
		KeySchema: []*dynamodb.KeySchemaElement{{
			AttributeName: aws.String("email"),
			KeyType:       aws.String(dynamodb.KeyTypeHash),
		}},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(25),
			WriteCapacityUnits: aws.Int64(25),
		},
		TableName: aws.String(db.UsersTableName()),
	})

	logrus.Info("Created Users table")

	hasCreatedUsers = err == nil
	return err
}

func (db *DB) createApplicationsTable() error {
	if hasCreatedApplications {
		return nil
	}

	err := CreateDynamoTableIfNotExists(db.client, &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{{
			AttributeName: aws.String("name"),
			AttributeType: aws.String(dynamodb.ScalarAttributeTypeS),
		}},
		KeySchema: []*dynamodb.KeySchemaElement{{
			AttributeName: aws.String("name"),
			KeyType:       aws.String(dynamodb.KeyTypeHash),
		}},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(25),
			WriteCapacityUnits: aws.Int64(25),
		},
		TableName: aws.String(db.ApplicationsTableName()),
	})

	logrus.Info("Created Applications table")

	hasCreatedApplications = err == nil
	return err
}

func (db *DB) createEnvironmentsTable() error {
	if hasCreatedEnvironments {
		return nil
	}

	err := CreateDynamoTableIfNotExists(db.client, &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{{
			AttributeName: aws.String("name"),
			AttributeType: aws.String(dynamodb.ScalarAttributeTypeS),
		}},
		KeySchema: []*dynamodb.KeySchemaElement{{
			AttributeName: aws.String("name"),
			KeyType:       aws.String(dynamodb.KeyTypeHash),
		}},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(25),
			WriteCapacityUnits: aws.Int64(25),
		},
		TableName: aws.String(db.EnvironmentsTableName()),
	})

	logrus.Info("Created Environments table")

	hasCreatedEnvironments = err == nil
	return err
}

func (db *DB) createDeploymentsTable() error {
	if hasCreatedDeployments {
		return nil
	}

	err := CreateDynamoTableIfNotExists(db.client, &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("app_env_version"),
				AttributeType: aws.String(dynamodb.ScalarAttributeTypeS),
			},{
				AttributeName: aws.String("app_env"),
				AttributeType: aws.String(dynamodb.ScalarAttributeTypeS),
			},{
				AttributeName: aws.String("timestamp"),
				AttributeType: aws.String(dynamodb.ScalarAttributeTypeS),
			},{
				AttributeName: aws.String("env_id"),
				AttributeType: aws.String(dynamodb.ScalarAttributeTypeS),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{{
			AttributeName: aws.String("app_env_version"),
			KeyType:       aws.String(dynamodb.KeyTypeHash),
		}},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(25),
			WriteCapacityUnits: aws.Int64(25),
		},
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{{
			IndexName: aws.String("app_env_timestamp"),
			KeySchema: []*dynamodb.KeySchemaElement{{
				AttributeName: aws.String("app_env"),
				KeyType:       aws.String(dynamodb.KeyTypeHash),
			}, {
				AttributeName: aws.String("timestamp"),
				KeyType:       aws.String(dynamodb.KeyTypeRange),
			}},
			Projection: &dynamodb.Projection{
				ProjectionType:   aws.String(dynamodb.ProjectionTypeAll),
			},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(25),
				WriteCapacityUnits: aws.Int64(25),
			},
		},{
			IndexName: aws.String("env_deploys"),
			KeySchema: []*dynamodb.KeySchemaElement{{
				AttributeName: aws.String("env_id"),
				KeyType:       aws.String(dynamodb.KeyTypeHash),
			},{
				AttributeName: aws.String("timestamp"),
				KeyType:       aws.String(dynamodb.KeyTypeRange),
			}},
			Projection: &dynamodb.Projection{
				ProjectionType:   aws.String(dynamodb.ProjectionTypeAll),
			},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(25),
				WriteCapacityUnits: aws.Int64(25),
			},
		}},
		TableName: aws.String(db.DeploymentsTableName()),
	})

	logrus.Info("Created Deployments table")

	hasCreatedDeployments = err == nil
	return err
}
