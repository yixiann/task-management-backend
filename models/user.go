package models

type User struct {
	Id        int64  `db:"idusers" json:"id"`
	Username  string `db:"username" json:"username"`
	Password  string `db:"password" json:"password"`
	Firstname string `db:"first_name" json:"firstName"`
	Lastname  string `db:"last_name" json:"lastName"`
}
