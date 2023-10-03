package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id          primitive.ObjectID   `bson:"_id, omniempty" json:id`
	Name        string               `json:name`
	Telephone   string               `json:telephone`
	Email       string               `json:email`
	Description string               `json:description`
	Tags        []string             `json:tags`
	Favorites   []primitive.ObjectID `json:favorites`
	Password    string               `json:password`
}

type UserLogin struct {
	Telephone string `json:telephone`
	Password  string `json:password`
	WebToken  string `json:webtoken`
}
