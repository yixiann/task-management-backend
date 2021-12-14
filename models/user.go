package models

type User struct {
	Id       int64  `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"password"`
	TaskId   string `db:"task_id" json:"taskId"`
}
