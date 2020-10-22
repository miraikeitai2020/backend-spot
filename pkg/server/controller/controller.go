package controller

import(
	"fmt"
	"net/http"

	// import gin library
	"github.com/gin-gonic/gin"

	// Import SQL driver
	"github.com/jinzhu/gorm"

	"github.com/miraikeitai2020/backend-spot/pkg/database"
	"github.com/miraikeitai2020/backend-spot/pkg/server/view"
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
		fmt.Printf("ERROR: %v\n", err)
		view.ReturnErrorResponse(
			cxt,
			http.StatusInternalServerError,
			"Internal Server Error",
			"Faild bind request JSON",
		)
	}

	spots := database.ElectSpotInfo(ctrl.DB)
	if len(spots) < 1 {
		fmt.Printf("ERROR: match infomation is empty.\n")
		view.ReturnErrorResponse(
			cxt,
			http.StatusInternalServerError,
			"Internal Server Error",
			"Spot is none in DB",
		)
	}

	view.ReturnGetSpotResponse(cxt, request, spots)
} 

func (ctrl *Controller) GetDetourHandler(cxt *gin.Context) {
	var request model.GetDetourRequest
	err := cxt.BindJSON(&request)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		view.ReturnErrorResponse(
			cxt,
			http.StatusInternalServerError,
			"Internal Server Error",
			"Faild bind request JSON",
		)
	}

	detour := database.ElectDetourInfo(ctrl.DB)
	if len(detour) < 1 {
		fmt.Printf("ERROR: match infomation is empty.\n")
		view.ReturnErrorResponse(
			cxt,
			http.StatusInternalServerError,
			"Internal Server Error",
			"Spot is none in DB",
		)
	}

	view.ReturnGetDetourResponse(cxt, request, detour)
}

func (ctrl *Controller) AddSpotHandler(cxt *gin.Context) {
	var request model.AddSpotRequest
	err := cxt.BindJSON(&request)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		view.ReturnErrorResponse(
			cxt,
			http.StatusInternalServerError,
			"Internal Server Error",
			"Faild bind request JSON",
		)
	}

	database.InsertDetourInfoToDataBase(ctrl.DB, request)
	view.ReturnMutationResponse(cxt, true)
}
