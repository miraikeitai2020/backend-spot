package view

import(
	"github.com/miraikeitai2020/backend-spot/pkg/server/model"
)

/*--------------------------------------*/
/*      Handler Response Generator      */
/*--------------------------------------*/
type GetSpotResponse struct {
	Spot	model.Spot	`json:"spot"`
}

func MakeGetSpotResponse(spot model.Spot) GetSpotResponse {
	return GetSpotResponse{
		Spot: spot,
	}
}

type GetDetourResponse struct {
	Detours	[]model.Detour	`json:"detours"`
}

func MakeGetDetourResponse(detour []model.Detour) GetDetourResponse {
	return GetDetourResponse{
		Detours: detour,
	}
}

type MutationResponse struct {
	Status	bool	`json:"status"`
}

func MakeMutationResponse(stat bool) MutationResponse {
	return MutationResponse{
		Status: stat,
	}
}

type Error struct {
	Code		int		`json:"code"`
	Message		string	`json:"message"`
	Description	string	`json:"description"`
}

func MakeErrorResponse(code int, msg, desc string) Error {
	return Error{
		Code: code,
		Message: msg,
		Description: desc,
	}
}
