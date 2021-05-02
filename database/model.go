package database

import "time"

// Model ...
type Model struct {
	ID        PID        `json:"id" gorm:"column:id;primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}

func (m Model) IsDeleted() bool {
	return m.DeletedAt != nil
}
