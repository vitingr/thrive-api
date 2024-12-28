package models

type User struct {
	ID                uint   `json:"id"`
	Username          string `json:"username"`
	Firstname         string `json:"firstname"`
	Lastname          string `json:"lastname"`
	Email             string `json:"email"`
	ProfilePicture    string `json:"profile_picture"`
	BackgroundPicture string `json:"background_picture"`
	Followers         int    `json:"followers"`
	Following         int    `json:"following"`
	Locale            string `json:"locale"`
	GoogleID          string `json:"google_id,omitempty"`
	Password          string `json:"password,omitempty"`
	Bio               string `json:"bio,omitempty"`

	Posts         []Post     `gorm:"foreignKey:CreatorId"`
	Likes         []Like     `gorm:"foreignKey:UserId"`
	FollowersList []Follower `gorm:"foreignKey:FollowingId"`
	FollowingList []Follower `gorm:"foreignKey:FollowerId"`
}

var Users []User
