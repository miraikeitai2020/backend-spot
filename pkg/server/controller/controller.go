package controller

import(
	"net/http"

	// import gin library
	"github.com/gin-gonic/gin"

	// Import SQL driver
	"github.com/jinzhu/gorm"

	"github.com/miraikeitai2020/backend-spot/pkg/spot"
	"github.com/miraikeitai2020/backend-spot/pkg/utils"
	"github.com/miraikeitai2020/backend-spot/pkg/database"
	"github.com/miraikeitai2020/backend-spot/pkg/server/model"
)

type Controller struct {
	DB *gorm.DB
}

func Init(db *gorm.DB) Controller {
	return Controller{
		DB: db,
	}
}

func (ctrl *Controller) GetSpotHandler(cxt *gin.Context) {
	var request model.GetSpotRequest
	err := cxt.BindJSON(&request)
	if err != nil {
		utils.LogFatal(err)
		cxt.JSON(
			http.StatusInternalServerError,
			utils.MakeErrorResponse(
				http.StatusInternalServerError,
				"Internal Server Error",
				"Faild bind request JSON",
			),
		)
	}

	// ここ脳死実装
	spots := []model.SpotInfo{}
	ctrl.DB.Raw(model.QUERY_FORMAT_GET_SPOT).Scan(&spots)
	if len(spots) == 0 {
		cxt.JSON(
			http.StatusInternalServerError,
			utils.MakeErrorResponse(
				http.StatusInternalServerError,
				"Internal Server Error",
				"Spot is none in DB",
			),
		)
	}

	cxt.JSON(
		http.StatusOK,
		utils.MakeGetSpotResponse(
			spot.Election(request, spots),
		),
	)
} 

func (ctrl *Controller) GetDetourHandler(cxt *gin.Context) {
	var request model.GetDetourRequest
	err := cxt.BindJSON(&request)
	if err != nil {
		utils.LogFatal(err)
		cxt.JSON(
			http.StatusInternalServerError,
			utils.MakeErrorResponse(
				http.StatusInternalServerError,
				"Internal Server Error",
				"Faild bind request JSON",
			),
		)
	}

	// TODO: Exec SQL query
	detour := model.Detour{}

	cxt.JSON(
		http.StatusOK,
		utils.MakeGetDetourResponse(detour),
	)
}

func (ctrl *Controller) AddSpotHandler(cxt *gin.Context) {
	var request model.AddSpotRequest
	err := cxt.BindJSON(&request)
	if err != nil {
		utils.LogFatal(err)
		cxt.JSON(
			http.StatusInternalServerError,
			utils.MakeErrorResponse(
				http.StatusInternalServerError,
				"Internal Server Error",
				"Faild bind request JSON",
			),
		)
	}

	// TODO: Exec SQL query
	database.InsertDetourInfoToDataBase(ctrl.DB, request)
	cxt.JSON(
		http.StatusOK,
		utils.MakeMutationResponse(true),
	)
}
