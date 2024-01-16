package models

// type IgPost struct {
// 	MediaUrl string   `json:mediaUrl`
// 	Caption  string   `json:Caption`
// 	Children []IgPost `json:children`
// }

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Publisher   primitive.ObjectID `json:"publisher" bson:"publisher"`
	Caption     string             `json:"caption" bson:"caption"`
	Comments    []Comment          `json:"comments" bson:"comments"`
	Likes       []primitive.ObjectID `json:"likes" bson:"likes"`
	Media       []string           `json:"media" bson:"media"`
	Timestamp   time.Time          `json:"timestamp" bson:"timestamp"`
}

type Comment struct {
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	Text      string             `json:"text" bson:"text"`
	Timestamp time.Time          `json:"timestamp" bson:"timestamp"`
}
