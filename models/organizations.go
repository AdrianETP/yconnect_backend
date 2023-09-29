package models

type Organization struct {
	Id          string   `json:_id`
	Name        string   `json:name`
	Location    string   `json:location`
	Description string   `json:description`
	Tags        []string `json:tags`
	Igtag       string   `json:igtag`
	FbTag       string   `json:fbtag`
	Images      string   `json:images`
	Telephone   string   `json:telephone`
	Email       string   `json:emaiL`
}
