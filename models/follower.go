package models

type Follower struct {
	Id          int `json:"id"`
	FollowerId  int `json:"follower_id"`
	FollowingId int `json:"following_id"`

	Follower  User `gorm:"foreignKey:FollowerId"`
	Following User `gorm:"foreignKey:FollowingId"`
}

var Followers []Follower
