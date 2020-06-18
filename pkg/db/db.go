package db

import (
	"github.com/ricoberger/go-vue-starter/pkg/db/dynamodb"
	"github.com/ricoberger/go-vue-starter/pkg/db/mongodb"
	"github.com/ricoberger/go-vue-starter/pkg/model"
	"github.com/sirupsen/logrus"
)

// Config represents the configuration of the database interface
type Config struct {
	MongoDB *mongodb.Config
	DynamoDB *dynamodb.Config
}

// DB is the interface which must be implemented by all db drivers
type DB interface {
	CloseConnection() error

	// User Management
	CreateUser(u *model.User) error
	GetUser(id string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	SaveUser(u *model.User) error
	DeleteUser(id string) error

	// Application management
	CreateApplication(app *model.Application) error
	SaveApplication(app *model.Application) error
	GetApplications() ([]*model.Application, error)
	BulkGetApplications(appIDs []string) ([]*model.Application, error)
	GetApplication(id string) (*model.Application, error)
	GetApplicationByName(name string) (*model.Application, error)

	// Environment management
	CreateEnvironment(env *model.Environment) error
	GetEnvironments() ([]*model.Environment, error)
	GetEnvironmentsForApp(appID string) ([]*model.Environment, error)
	GetEnvironmentByName(name string) (*model.Environment, error)

	// Deployment management
	CreateDeployment(deploy *model.Deployment) error
	GetDeployment(id string) (*model.Deployment, error)
	GetDeployments() ([]*model.Deployment, error)
	GetLatestDeployForAppEnv(appID string, envID string) (*model.Deployment, error)
	GetDeploymentsForApp(appID string) ([]*model.Deployment, error)
	GetDeploymentsForEnv(envID string) ([]*model.Deployment, error)
}

// NewConnection creates a new database connection
func NewConnection(config *Config) (DB, error) {
	// Use MongoDB
	/*db, err := mongodb.NewConnection(config.MongoDB)
	if err != nil {
		return nil, err
	}*/

	db, err := dynamodb.New(config.DynamoDB)
	if err != nil {
		return nil, err
	}

	if config.DynamoDB.Env == "local" {
		logrus.Info("Running locally, creating tables...")
		err = db.SeedTables()
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
