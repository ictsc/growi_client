package entity

import "time"

type Page struct {
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
	LastUpdateUser  LastUpdateUser `json:"lastUpdateUser"`
	GrantedGroup    interface{}    `json:"grantedGroup"`
	V               int            `json:"__v"`
	Revision        Revision       `json:"revision"`
}

type Revision struct {
	Id     string `json:"_id"`
	PageId string `json:"pageId"`
	Body   string `json:"body"`
	V      int    `json:"__v"`
}
