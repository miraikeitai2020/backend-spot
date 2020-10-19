package server

import(
	// import gin library
	"github.com/gin-gonic/gin"

	"github.com/miraikeitai2020/backend-spot/pkg/server/controller"
)

func Router(ctrl controller.Controller) *gin.Engine {
	router := gin.Default()
	// #3 スポット選出のためのハンドラ作成
	router.POST("/query/spots", ctrl.GetSpotsHandler)
	// #4 寄り道スポット選出のためのハンドラ作成
	router.POST("/query/detour", ctrl.GetDetourHandler)
	// #5 スポット登録のハンドラ作成
	router.POST("/mutation/add/spot", ctrl.AddSpotHandler)
	return router
}
