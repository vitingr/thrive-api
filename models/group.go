package models

type Group struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	Activities        string `json:"activities"`
	GroupPicture      string `json:"group_picture"`
	BackgroundPicture string `json:"background_picture"`
	IsPrivate         bool   `json:"is_private"`
	Followers         int    `json:"followers"`
	Locale            string `json:"locale"`
	Members           int    `json:"members"`
}

var Groups []Group
