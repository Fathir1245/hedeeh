package model

import "time"

// BaseModel berisi field umum yang hampir selalu ada
// Bisa di-embed ke model lain
type BaseModel struct {
	ID        uint      `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
