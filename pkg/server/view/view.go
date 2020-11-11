package view

import(
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/miraikeitai2020/backend-spot/pkg/spot"
	"github.com/miraikeitai2020/backend-spot/pkg/server/model"
)

/*--------------------------------------*/
/*      Handler Response Generator      */
/*--------------------------------------*/
type GetSpotResponse struct {
	Spot	model.Spot	`json:"spot"`
}

func ReturnGetSpotResponse(cxt *gin.Context, request model.GetSpotRequest, spots []model.SpotInfo) {
	body := GetSpotResponse{
		Spot: spot.Election(request, spots),
	}
	cxt.JSON(http.StatusOK, body)
}

type GetDetourResponse struct {
	Detours	[]model.Detour	`json:"detours"`
}

func ReturnGetDetourResponse(cxt *gin.Context, request model.GetDetourRequest, detour []model.DetourInfo) {
	body := GetDetourResponse{
		Detours: spot.DetourElection(request, detour),
	}
	cxt.JSON(http.StatusOK, body)
}

type MutationResponse struct {
	Status	bool	`json:"status"`
}

func ReturnMutationResponse(cxt *gin.Context, stat bool) {
	body := MutationResponse{
		Status: stat,
	}
	cxt.JSON(http.StatusOK, body)
}

type Error struct {
	Code		int		`json:"code"`
	Message		string	`json:"message"`
	Description	string	`json:"description"`
}

func ReturnErrorResponse(cxt *gin.Context, code int, msg, desc string) {
	body := Error{
		Code: code,
		Message: msg,
		Description: desc,
	}
	cxt.JSON(code, body)
}
