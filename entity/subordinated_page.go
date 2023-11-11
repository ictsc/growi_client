package entity

import "time"

type SubordinatedPage struct {
	Id              string         `json:"_id"`
	Parent          string         `json:"parent"`
	DescendantCount int            `json:"descendantCount"`
	IsEmpty         bool           `json:"isEmpty"`
	Status          string         `json:"status"`
	Grant           int            `json:"grant"`
	GrantedUsers    []interface{}  `json:"grantedUsers"`
	Liker           []interface{}  `json:"liker"`
	SeenUsers       []string       `json:"seenUsers"`
	CommentCount    int            `json:"commentCount"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	Path            string         `json:"path"`
	Creator         string         `json:"creator"`
	LastUpdateUser  LastUpdateUser `json:"lastUpdateUser"`
	GrantedGroup    interface{}    `json:"grantedGroup"`
	V               int            `json:"__v"`
	Revision        string         `json:"revision"`
}
