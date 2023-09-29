package models

type User struct {
    Id string `json:_id`
	Name        string   `json:name`
	Telephone   string   `json:telephone`
	Email       string   `json:email`
	Description string   `json:description`
	Tags        []string `json:tags`
	Favorites   []string `json:favorites`
	Password    string   `json:password`
}

type UserLogin struct {
	Telephone string `json:telephone`
	Password  string `json:password`
	WebToken  string `json:webtoken`
}
