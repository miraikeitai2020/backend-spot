package spot

import(
	"math"
	"sort"
	"github.com/miraikeitai2020/backend-spot/pkg/server/model"
)

// 地球の長半径
const rx = 6378137.0

type info struct {
	Latitude	float64
	Longitude	float64
}

func deg2rad(x float64) float64 {
	return x * math.Pi / 180.0
}

func calcUserDistance(uLat, uLon, sLat, sLon float64) float64 {
	x0 := deg2rad(uLat)
	y0 := deg2rad(uLon)
	x1 := deg2rad(sLat)
	y1 := deg2rad(sLon)

	return rx * math.Acos(math.Sin(y0)*math.Sin(y1)+math.Cos(y0)*math.Cos(y1)*math.Cos(x0 - x1))
}

func isWithinRange(time float64, lat, lon float64, s info) bool {
	return calcUserDistance(lat, lon, s.Latitude, s.Longitude) / 80.0 < time
}

func spotCandidates(time float64, lat, lon float64, s []model.SpotInfo) (spots []model.SpotInfo) {
	for _, v := range s {
		i := info{
			Latitude: v.Latitude,
			Longitude: v.Longitude,
		}
		if isWithinRange(time, lat, lon, i) {
			spots = append(spots, v)
		}
	}
	return
}

func detourCandidates(time float64, lat, lon float64, s []model.DetourInfo) (detour []model.DetourInfo) {
	for _, v := range s {
		i := info{
			Latitude: v.Latitude,
			Longitude: v.Longitude,
		}
		if isWithinRange(time, lat, lon, i) {
			detour = append(detour, v)
		}
	}
	return
} 

func comparisonValue(emotion int, _v, v model.SpotInfo) model.SpotInfo {
	switch emotion {
	case 0:
		if _v.Happiness < v.Happiness {
			_v = v
		}
	case 1:
		if _v.Calm < v.Calm {
			_v = v
		}
	case 2:
		if _v.Exciting < v.Exciting {
			_v = v
		}
	case 3:
		if _v.Sadness < v.Sadness {
			_v = v
		}
	}
	return _v
}
func detourcomparisonValue(emotion int, v []model.DetourInfo) []model.DetourInfo {
	switch emotion {
	case 0:
		
		sort.Slice(v,func(i,j int)bool{return v[i].Happiness<v[j].Happiness})
	
	case 1:
		
		sort.Slice(v,func(i,j int)bool{return v[i].Calm<v[j].Calm})

	case 2:
	
		sort.Slice(v,func(i,j int)bool{return v[i].Exciting<v[j].Exciting})

	case 3:
		
		sort.Slice(v,func(i,j int)bool{return v[i].Sadness<v[j].Sadness})

	
	}
	return v
}
func emotionFilter(emotion int, s []model.SpotInfo) model.Spot {
	electionSpotInfo := s[0]	//s０番目を格納
	for _, v := range s {
		electionSpotInfo = comparisonValue(emotion, electionSpotInfo, v)	
	}
	return model.Spot{
		ID: electionSpotInfo.ID,
		Name: electionSpotInfo.Name,
		Latitude: electionSpotInfo.Latitude,
		Longitude: electionSpotInfo.Longitude,
	}
}
func detouremotionFilter(emotion int, s []model.DetourInfo) []model.Detour {
	
	var Detours []model.Detour
	
		s= detourcomparisonValue(emotion, s)//ソートをかける
	for _,v:=range s{
		electionDetourInfo:=model.Detour{
			ID: v.ID,
			Name: v.Name,
			Image: v.Image,
			Latitude: v.Latitude,
			Longitude: v.Longitude,
		}
		Detours=append(Detours,electionDetourInfo)

	}
	
	
	return Detours
}
func Election(r model.GetSpotRequest, s []model.SpotInfo) model.Spot {
	spots := spotCandidates(float64(r.Walktime)*0.9, r.Latitude, r.Longitude, s)
	return emotionFilter(r.Emotion, spots)
}

func DetourElection(r model.GetDetourRequest, s []model.DetourInfo) []model.Detour {
	
	detours :=  detourCandidates(float64(r.Walktime)*0.45, (r.UserLatitude + r.SpotLatitude)/2.0, (r.UserLongitude + r.SpotLongitude)/2.0, s)
	return detouremotionFilter(r.Emotion,detours)
}
