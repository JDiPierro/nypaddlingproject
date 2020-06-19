package model

import "time"

type Location struct {
	ID                string    `json:"id" bson:"_id"`
	Title             string    `json:"title" bson:"title"`
	County            string    `json:"county" bson:"county"`
	CreatedAt         time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" bson:"updated_at"`
	DescriptionLength int       `json:"desc_len" bson:"desc_len"`
	NumPhotos         int       `json:"num_photos" bson:"num_photos"`
	Link              string    `json:"link" bson:"link"`
}
