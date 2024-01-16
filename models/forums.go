package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Forums struct {
	Id        primitive.ObjectID `bson:"_id, omniempty" json:id`
	UserId    primitive.ObjectID `bson:"userId, omniempty" json:userId`
	OrgId     primitive.ObjectID `bson:"orgId, omniempty" json:orgId`
	Content   string             `bson:"content, omniempty" json:content`
	TimeStamp primitive.DateTime `bson:"timestamp, omniempty" json:timestamp`
}
