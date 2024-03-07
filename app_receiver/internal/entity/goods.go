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
	ID          int64     `db:"Id"`
	ProjectID   int64     `db:"ProjectID"`
	Name        string    `db:"Name"`
	Description string    `db:"Description"`
	Priority    int64     `db:"Priority"`
	Removed     bool      `db:"Removed"`
	EventTime   time.Time `db:"EventTime"`
}
