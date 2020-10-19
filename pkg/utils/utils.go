package utils

import(
	"log"
	"github.com/miraikeitai2020/backend-spot/pkg/server/model"
)

func LogFatal(err error) {
	log.Fatal(err)
}

func MakeErrorResponse(code int, msg, desc string) model.Error {
	return model.Error{
		Code: code,
		Message: msg,
		Description: desc,
	}
}

func MakeGetSpotsResponse(spots []model.Spot) model.GetSpotsResponse {
	return model.GetSpotsResponse{
		Spots: spots,
	}
}

func MakeGetDetourResponse(detour model.Detour) model.GetDetourResponse {
	return model.GetDetourResponse{
		Detour: detour,
	}
}
