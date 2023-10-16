package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Organization struct {
	ID          primitive.ObjectID `bson:"_id, omitempty" json:id`
	Name        string             `json:name`
	Location    string             `json:location`
	Description string             `json:description`
	Tags        []string           `json:tags`
	Igtag       string             `json:igtag`
	IgUrl       string             `json:igurl`
	FbTag       string             `json:fbtag`
	Images      string             `json:images`
	Telephone   string             `json:telephone`
	Email       string             `json:emaiL`
}
