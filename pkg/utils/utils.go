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

func MakeGetSpotResponse(spot model.Spot) model.GetSpotResponse {
	return model.GetSpotResponse{
		Spot: spot,
	}
}

func MakeGetDetourResponse(detour model.Detour) model.GetDetourResponse {
	return model.GetDetourResponse{
		Detour: detour,
	}
}

func MakeMutationResponse(stat bool) model.MutationResponse {
	return model.MutationResponse{
		Status: stat,
	}
}
