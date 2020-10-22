package database

import(
	"fmt"

	// import gorm library
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	
	"github.com/miraikeitai2020/backend-spot/pkg/spot"
	"github.com/miraikeitai2020/backend-spot/pkg/uuid"
	"github.com/miraikeitai2020/backend-spot/pkg/utils"
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

func InsertDetourInfoToDataBase(db *gorm.DB, r model.AddSpotRequest) {
	id := uuid.GenerateUUID()
	db.Exec(model.QUERY_FORMAT_ADD_DETOURS, id, r.Name, r.Description, r.Latitude, r.Longitude)
	db.Exec(model.QUERY_FORMAT_ADD_DETOURS_EMOTION, id, utils.Random(0.0, 0.3), utils.Random(0.0, 0.3), utils.Random(0.0, 0.3), utils.Random(0.0, 0.3))
	db.Exec(model.QUERY_FORMAT_ADD_DETOURS_IMAGE, id, []byte(r.Image))
}

func ElectDetourInfo(db *gorm.DB, r model.GetDetourRequest) []model.Detour {
	info := []model.Detour{}
	db.Raw(model.QUERY_FORMAT_GET_DETOURS).Scan(&info)
	return spot.DetourElection(r, info)
}
