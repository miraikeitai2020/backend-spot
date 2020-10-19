package controller

import(
	// Import SQL driver
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Controller struct {
	DB *gorm.DB
}

func Init(db *gorm.DB) Controller {
	return Controller{
		DB: db,
	}
}
