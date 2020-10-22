package spot

import(
	"math"

	"github.com/miraikeitai2020/backend-spot/pkg/server/model"
)

func deg2rad(x float64) float64 {
	return x * math.Pi / 180.0
}

func calcUserDistance(uLat, uLon, sLat, sLon float64) float64 {
	rx := 6378137.0

	x0 := deg2rad(uLat)
	y0 := deg2rad(uLon)
	x1 := deg2rad(sLat)
	y1 := deg2rad(sLon)

	return rx * math.Acos(math.Sin(y0)*math.Sin(y1)+math.Cos(y0)*math.Cos(y1)*math.Cos(x0 - x1))
}

func isWithinRange(time int, lat, lon float64, s model.SpotInfo) bool {
	return calcUserDistance(lat, lon, s.Latitude, s.Longitude) / 80.0 < float64(time)
}

func spotCandidates(time int, lat, lon float64, s []model.SpotInfo) (spots []model.SpotInfo) {
	for _, v := range s {
		if isWithinRange(time, lat, lon, v) {
			spots = append(spots, v)
		}
	}
	return
}

func emotionFilter(emotion int, s []model.SpotInfo) model.Spot {
	_s := model.SpotInfo{
		Happiness: 0.0,
		Natural: 0.0,
		Sadness: 0.0,
		Anger: 0.0,
	}
	// ここ改善よちアリ
	// _s = _v の時のアルゴリズム
	// P.S. 高速化できなかなぁ
	for _, v := range s {
		switch emotion {
		case 1:
			if _s.Happiness < v.Happiness {
				_s = v
			}
		case 2:
			if _s.Natural < v.Natural {
				_s = v
			}
		case 3:
			if _s.Sadness < v.Sadness {
				_s = v
			}
		case 4:
			if _s.Anger < v.Anger {
				_s = v
			}
		}
	}
	return model.Spot{
		ID: _s.ID,
		Name: _s.Name,
		Latitude: _s.Latitude,
		Longitude: _s.Longitude,
	}
}

func Election(r model.GetSpotRequest, s []model.SpotInfo) model.Spot {
	spots := spotCandidates(r.Walktime, r.Latitude, r.Longitude, s)
	return emotionFilter(r.Emotion, spots)
}
