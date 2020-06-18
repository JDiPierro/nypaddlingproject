package model

// Deployment represents a version of an Application in an Environment
type Deployment struct {
	ID        string `json:"id" bson:"_id"`
	AppID     string `json:"app_id" bson:"app_id"`
	EnvID     string `json:"env_id" bson:"env_id"`
	Version   string `json:"version" bson:"version"`
	User      string `json:"user" bson:"user"`
	Branch    string `json:"branch" bson:"branch"`
	Message   string `json:"message" bson:"message"`
	Timestamp string `json:"timestamp" bson:"timestamp"`
}
