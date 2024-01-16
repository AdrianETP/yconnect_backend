package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Forums struct {
	Id          primitive.ObjectID   `bson:"_id, omniempty" json:id`
	User     	primitive.ObjectID   `bson:"user, omniempty" json:user`
	Org    		primitive.ObjectID   `bson:"org, omniempty" json:org`
	Content 	string `bson:"content, omniempty" json:content`
	TimeStamp 	string `bson:"timestamp, omniempty" json:timestamp`
}
