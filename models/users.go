package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id          primitive.ObjectID   `bson:"_id, omniempty" json:id`
	Name        string               `json:name`
	Telephone   string               `json:telephone`
	Description string               `json:description`
	Tags        []string             `json:tags`
	Favorites   []primitive.ObjectID `json:favorites`
	Password    string               `json:password`
	Email       string               `json:email`
}

type UserLogin struct {
	Email    string `json:email`
	Password string `json:password`
}
