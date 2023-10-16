package models

type IgPost struct {
	MediaUrl string   `json:mediaUrl`
	Caption  string   `json:Caption`
	Children []IgPost `json:children`
}
