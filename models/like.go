package models

type Like struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
	PostId int `json:"post_id"`

	User User `gorm:"foreignKey:UserId"`
	Post Post `gorm:"foreignKey:PostId"`
}

var Likes []Like
