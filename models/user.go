package models

type User struct {
	Uid               string  `json:"uid"`
	Username          string  `json:"username"`
	Firstname         string  `json:"firstname"`
	Lastname          string  `json:"lastname"`
	Email             string  `json:"email"`
	ProfilePicture    string  `json:"profile_picture"`
	BackgroundPicture string  `json:"background_picture"`
	Followers         int     `json:"followers"`
	Following         int     `json:"following"`
	GoogleID          *string `json:"google_id,omitempty"`
}

var Users []User
