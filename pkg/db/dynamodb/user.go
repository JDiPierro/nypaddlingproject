package dynamodb

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/ricoberger/go-vue-starter/pkg/model"
)

// CreateUser creates a new user
func (db *DB) CreateUser(u *model.User) error {
	u.ID = u.Email
	return db.SaveUser(u)
}

// GetUser returns a user
func (db *DB) GetUser(id string) (*model.User, error) {
	return db.GetUserByEmail(id)
}

// GetUserByEmail returns a user by his email address
func (db *DB) GetUserByEmail(email string) (*model.User, error) {
	output, err := db.client.GetItem(&dynamodb.GetItemInput{
		ConsistentRead: aws.Bool(true),
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
		},
		TableName: aws.String(db.UsersTableName()),
	})

	if err != nil {
		return nil, err
	} else if output == nil || output.Item == nil {
		return nil, fmt.Errorf("no user found with the email: %s", email)
	}

	result := &model.User{}
	err = dynamodbattribute.UnmarshalMap(output.Item, result)
	return result, err
}

// SaveUser saves the given user struct
func (db *DB) SaveUser(u *model.User) error {
	item, err := dynamodbattribute.MarshalMap(u)
	if err != nil {
		return err
	}

	_, err = db.client.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(db.UsersTableName()),
	})

	return err
}

// DeleteUser deletes the user with the given id
func (db *DB) DeleteUser(id string) error {
	_, _ = db.client.DeleteItem(&dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(db.UsersTableName()),
	})
	return nil
}
