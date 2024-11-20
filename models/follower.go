package models

type Follower struct {
	Id          int `json:"id"`
	FollowerId  int `json:"follower_id"`
	FollowingId int `json:"following_id"`
}

var Followers []Follower
