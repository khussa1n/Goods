package entity

import "time"

type GoodsResponce struct {
	ID          int64     `json:"id"`
	ProjectID   int64     `json:"project_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Priority    int64     `json:"priority"`
	Removed     bool      `json:"removed"`
	CreatedAt   time.Time `json:"created_at"`
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
