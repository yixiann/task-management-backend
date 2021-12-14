package models

type Task struct {
	Id         int64  `db:"id" json:"id"`
	UserId     int    `db:"user_id" json:"userId"`
	TaskName   string `db:"task_name" json:"taskName"`
	Details    string `db:"details" json:"details"`
	TagId      string `db:"tag_id" json:"tagId"`
	Deadline   string `db:"deadline" json:"deadline"`
	Priority   string `db:"priority" json:"priority"`
	TaskStatus string `db:"task_status" json:"taskStatus"`
	CreatedBy  string `db:"created_by" json:"createdBy"`
	AssignedTo string `db:"assigned_to" json:"assignedTo"`
}
