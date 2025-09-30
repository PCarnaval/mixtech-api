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
	remote  bool
}

type OpeningResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Role      string    `json:""`
	Name      string    `json:""`
	Company   string    `json:""`
	Email     string    `json:""`
	Salary    int64     `json:""`
	remote    bool      `json:""`
}
