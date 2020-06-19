package model

type User struct {
	ID            string `json:"id" bson:"_id"`
	FirstName     string `json:"first_name" bson:"first_name"`
	LastName      string `json:"last_name" bson:"last_name"`
	Role          string `json:"role" bson:"role"`
	Points        int    `json:"points" bson:"points"`
	FBAccessToken string `json:"fb_access_token" bson:"fb_access_token"`
}
