package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Testimonio struct {
	Id        primitive.ObjectID `bson:"_id, omniempty" json:id`
	UserId    primitive.ObjectID `bson:"userId, omniempty" json:userId`
	Content   string             `bson:"content, omniempty" json:content`
	Title     string             `bson:"title, omniempty" json:title`
	OrgId     primitive.ObjectID `bson:"orgId, omniempty" json:orgId`
	TimeStamp primitive.DateTime `bson:"timestamp, omniempty" json:timestamp`
	Grade     int                `bson:"grade, omniempty" json:grade`
}
