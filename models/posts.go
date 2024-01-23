package models

// type IgPost struct {
// 	MediaUrl string   `json:mediaUrl`
// 	Caption  string   `json:Caption`
// 	Children []IgPost `json:children`
// }

import (

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id          primitive.ObjectID `bson:"_id, omniempty" json:id`
	OrgId   primitive.ObjectID `bson:"orgId, omniempty" json:orgId`
	Content     string             `bson:"content, omniempty" json:content`
	Comments    []Comment          `json:"comments" bson:"comments"`
	Likes       []primitive.ObjectID `bson:"likes, omniempty" json:likes`
	MediaUrls       []string           `bson:"media, omniempty" json:media`
	TimeStamp   primitive.DateTime          `bson:"timestamp, omniempty" json:timestamp`
}

type Comment struct {
	UserID    primitive.ObjectID `bson:"userId, omniempty" json:userId`
	Content      string             `bson:"content, omniempty" json:content`
	TimeStamp primitive.DateTime         `bson:"media, omniempty" json:media`
}
