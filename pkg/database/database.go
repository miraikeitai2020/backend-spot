package database

import(
	"fmt"

	// import gorm library
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	
	"github.com/miraikeitai2020/backend-spot/pkg/uuid"
	"github.com/miraikeitai2020/backend-spot/pkg/config"
	"github.com/miraikeitai2020/backend-spot/pkg/server/model"
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

func ElectSpotInfo(db *gorm.DB) (info []model.SpotInfo) {
	db.Raw(model.QUERY_FORMAT_GET_SPOT).Scan(&info)
	return
}

func ElectDetourInfo(db *gorm.DB) (info []model.DetourInfo) {
	db.Raw(model.QUERY_FORMAT_GET_DETOURS).Scan(&info)
	return
}

func InsertDetourInfoToDataBase(db *gorm.DB, r model.AddSpotRequest) {
	id := uuid.GenerateUUID()
	db.Exec(model.QUERY_FORMAT_ADD_DETOURS, id, r.Name, r.Description, r.Latitude, r.Longitude)
	db.Exec(model.QUERY_FORMAT_ADD_DETOURS_IMAGE, id, []byte(r.Image))
	switch r.Emotion {
	case 0:
		db.Exec(model.QUERY_FORMAT_ADD_DETOURS_EMOTIONS, id,0.6,0,0,0)
	case 1:
		db.Exec(model.QUERY_FORMAT_ADD_DETOURS_EMOTIONS, id,0,0.6,0,0)

	case 2:
		db.Exec(model.QUERY_FORMAT_ADD_DETOURS_EMOTIONS, id,0,0,0.6,0)

	case 3:
		db.Exec(model.QUERY_FORMAT_ADD_DETOURS_EMOTIONS, id,0,0,0,0.6)

	}
	
}
