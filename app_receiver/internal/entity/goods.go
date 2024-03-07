package entity

import "time"

type GoodsResponce struct {
	ID          int64     `json:"id" db:"id"`
	ProjectID   int64     `json:"project_id" db:"project_id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Priority    int64     `json:"priority" db:"priority"`
	Removed     bool      `json:"removed" db:"removed"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type Goods struct {
	ID          int64     `db:"id"`
	ProjectID   int64     `db:"project_id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Priority    int64     `db:"priority"`
	Removed     bool      `db:"removed"`
	EventTime   time.Time `db:"created_at"`
}
