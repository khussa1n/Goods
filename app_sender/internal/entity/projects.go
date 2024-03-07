package entity

import "time"

type Projects struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
