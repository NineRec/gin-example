package storage

import (
	"time"
)

// Todo ...
type Todo struct {
	ID          int64
	Author      string `gorm:"size:255" json:"author"`
	Description string `gorm:"size:255" json:"description"`
	Done        bool
	CreatedAt   time.Time
}
