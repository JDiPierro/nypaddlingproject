package model

// Environments are the target of a Deployment
type Environment struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}
