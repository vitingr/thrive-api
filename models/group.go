package models

type Group struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	FavouriteActivity string `json:"favourite_activity"`
	GroupPicture      string `json:"group_picture"`
	BackgroundPicture string `json:"background_picture"`
	IsPrivate         bool   `json:"is_private"`
	Followers         int    `json:"followers"`
	Locale            int    `json:"string"`
	Members           int    `json:"members"`
}

var Groups []Group
