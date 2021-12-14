package models

type Tag struct {
	Id      int    `json:"id"`
	TagName string `json:"tagName"`
	Colour  string `json:"colour"`
}
