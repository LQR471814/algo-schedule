// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"time"
)

type Project struct {
	ID          int64
	DeletedAt   sql.NullTime
	Name        string
	Description string
	Deadline    time.Time
}

type ProjectTask struct {
	ID          int64
	ProjectID   int64
	DeletedAt   sql.NullTime
	Name        string
	Description string
	Size        int64
}

type Quotum struct {
	ID                 int64
	DeletedAt          sql.NullTime
	Description        interface{}
	FixedTime          int64
	Duration           int64
	RecurrenceInterval int64
}

type Setting struct {
	ID       int64
	Timezone string
}

type Task struct {
	ID          int64
	DeletedAt   sql.NullTime
	Name        string
	Description string
	Deadline    time.Time
	Size        int64
	Priority    int64
}
