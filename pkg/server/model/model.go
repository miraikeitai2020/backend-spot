package model

type GetSpotsRequest struct {
	Latitude	float64	`json:"latitude"`
	Longitude	float64	`json:"longitude"`
	// spot_required
	// スポットへの所要時間
	Walktime	int		`json:"walktime"`
	Emotion		int		`json:"emotion"`
}

type GetSpotsResponse struct {
	Spots	[]Spot	`json:"spots"`
}

type Spot struct {
	ID			string	`json:"id"`
	// 保留カラム
	Name		string	`json:"name"`
	Latitude	float64	`json:"latitude"`
	Longitude	float64	`json:"longitude"`
}

type Error struct {
	Code		int		`json:"code"`
	Message		string	`json:"message"`
	Description	string	`json:"description"`
}
