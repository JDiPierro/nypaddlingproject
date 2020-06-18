package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Env          string `yaml:"env"`
	DynamoURL    string `yaml:"url"`
	DynamoRegion string `yaml:"region"`
}

type DB struct {
	client *dynamodb.DynamoDB
	cfg    *Config
}

// New creates a new instance of DB
func New(cfg *Config) (*DB, error) {
	awsCfg := aws.NewConfig()

	if cfg.Env == "local" {
		logrus.Debug("Setting static AWS credentials for local testing: ", cfg.DynamoURL)
		awsCfg.Endpoint = aws.String(cfg.DynamoURL)
		awsCfg.Credentials = credentials.NewStaticCredentials("test", "test", "")
	}
	awsCfg.Region = aws.String(cfg.DynamoRegion)

	sess, err := session.NewSession(awsCfg)
	if err != nil {
		return nil, errors.Wrap(err, "dynamo was unable to create a new session")
	}

	client := dynamodb.New(sess, awsCfg)

	return &DB{
		client: client,
		cfg:    cfg,
	}, nil
}

// CloseConnection closes the database connection
func (db *DB) CloseConnection() error {
	return nil
}
