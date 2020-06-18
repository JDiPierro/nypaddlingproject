package model

// Application gets deployed to an Environment
type Application struct {
	ID         string `json:"id" bson:"_id"`
	Name       string `json:"name" bson:"name"`
}
