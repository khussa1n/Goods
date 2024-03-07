package entity

import "time"

type Goods struct {
	ID          int64     `json:"id" db:"id"`
	ProjectID   int64     `json:"project_id" db:"project_id"`
	Name        string    `json:"name" db:"name" binding:"required"`
	Description string    `json:"description" db:"description" binding:"required"`
	Priority    int64     `json:"priority" db:"priority"`
	Removed     bool      `json:"removed" db:"removed"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
