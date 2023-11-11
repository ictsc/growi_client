package entity

import "time"

type LastUpdateUser struct {
	Id                string    `json:"_id"`
	IsGravatarEnabled bool      `json:"isGravatarEnabled"`
	IsEmailPublished  bool      `json:"isEmailPublished"`
	Lang              string    `json:"lang"`
	Status            int       `json:"status"`
	Admin             bool      `json:"admin"`
	CreatedAt         time.Time `json:"createdAt"`
	Name              string    `json:"name"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	ImageUrlCached    string    `json:"imageUrlCached"`
	LastLoginAt       time.Time `json:"lastLoginAt"`
}
