package handler

import (
	"github.com/PCarnaval/mixtech-api/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
	_ = db // Use db to avoid unused variable error
}
