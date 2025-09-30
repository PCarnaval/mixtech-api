package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role    string
	Name    string
	Company string
	Email   string
	Salary  int64
	Remote  bool
}

type OpeningResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Role      string    `json:"role"`
	Name      string    `json:"name"`
	Company   string    `json:"company"`
	Email     string    `json:"email"`
	Salary    int64     `json:"salary"`
	Remote    bool      `json:"remote"`
}
