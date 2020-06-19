package model

import "time"

type Claim struct {
	ID         string    `json:"id" bson:"_id"`
	UserID     string    `json:"user_id" bson:"user_id"`
	LocationID string    `json:"location_id" bson:"location_id"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
	Comment    string    `json:"comment" bson:"comment"`
	Status     string    `json:"status" bson:"status"`
	Points     int       `json:"points" bson:"points"`
	Reason     string    `json:"reason" bson:"reason"`
}
