package schemas

import (
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
