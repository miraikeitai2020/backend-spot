package controller

import(
	"net/http"

	// import gin library
	"github.com/gin-gonic/gin"

	// Import SQL driver
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/miraikeitai2020/backend-spot/pkg/server/model"
	"github.com/miraikeitai2020/backend-spot/pkg/utils"
)

type Controller struct {
	DB *gorm.DB
}

func Init(db *gorm.DB) Controller {
	return Controller{
		DB: db,
	}
}

func (ctrl *Controller) GetSpotsHandler(cxt *gin.Context) {
	var request model.GetSpotsRequest
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
	spots := []model.Spot{}

	cxt.JSON(
		http.StatusOK,
		utils.MakeGetSpotsResponse(spots),
	)
} 

func (ctrl *Controller) GetDetoursHandler(cxt *gin.Context) {
}

func (ctrl *Controller) AddSpotHandler(cxt *gin.Context) {
} 
