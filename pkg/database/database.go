package database

import(
	"fmt"

	// import gorm library
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	
	"github.com/miraikeitai2020/backend-spot/pkg/config"
)

const(
	DB_DRIVER_MYSQL = "mysql"
)

func Init() (*gorm.DB, error) {
	token := config.GetConnectionToken()
	for i:=0; i<5; i++ {
		if db, err := gorm.Open(DB_DRIVER_MYSQL, token);err == nil {
			return db, nil
		}
		config.BackOff(i)
	}
	return nil, fmt.Errorf("Faild connection database")
}
