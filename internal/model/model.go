package model

import (
	"time"
)

type FolderWithTasks struct {
	Folder
	Tasks []Task
}

type Folder struct {
	ID   int       `json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

type NewFolder struct {
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

type Task struct {
	ID       int       `json:"id"`
	Name     string    `json:"name" db:"task_name"`
	Text     string    `json:"text" db:"task_text"`
	Date     time.Time `json:"date" db:"task_date"`
	Status   bool      `json:"status" db:"status"`
	FolderID int       `json:"folderID" db:"folders_id"`
}

type NewTask struct {
	Name     string    `db:"task_name"`
	Text     string    `db:"task_text"`
	Date     time.Time `db:"task_date"`
	Status   bool      `db:"status"`
	FolderID int       `db:"folders_id"`
}
