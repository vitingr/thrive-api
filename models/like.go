package models

type Like struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	PostId    int    `json:"post_id"`
}

var Likes []Like
