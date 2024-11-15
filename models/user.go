package models

type User struct {
	Username          string  `json:"username"`
	Firstname         string  `json:"firstname"`
	Lastname          string  `json:"lastname"`
	Email             string  `json:"email"`
	ProfilePicture    string  `json:"profile_picture"`
	BackgroundPicture string  `json:"background_picture"`
	Followers         int     `json:"followers"`
	Following         int     `json:"following"`
	Locale            string  `json:"locale"`
	GoogleID          *string `json:"google_id,omitempty"`
}

var Users []User
