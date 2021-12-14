package models

type Tag struct {
	Id      int64  `db:"id" json:"id"`
	TagName string `db:"tag_name" json:"tagName"`
	Colour  string `db:"colour" json:"colour"`
}
