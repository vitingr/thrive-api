package models

type Post struct {
	Id          int    `json:"id"`
	Content     string `json:"content"`
	Location    string `json:"location"`
	ImageUrl    string `json:"image_url"`
	VideoUrl    string `json:"video_url"`
	Type        string `json:"type"`
	CreatorId   int    `json:"creator_id"`
	NumberLikes int    `json:"number_likes"`
	Creator     User   `json:"creator" gorm:"foreignKey:CreatorId;references:Id"`
}

var Posts []Post
