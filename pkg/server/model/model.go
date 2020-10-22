package model


// SQL Query
const(
	QUERY_FORMAT_GET_SPOT = "SELECT spots.id, spots.name, spots.latitude, spots.longitude, emotions.happiness, emotions.natural, emotions.sadness, emotions.anger FROM `spots`, `emotions` WHERE spots.id = emotions.id GROUP BY spots.id"
	QUERY_FORMAT_GET_DETOURS = "SELECT * FROM `detours`"
	QUERY_FORMAT_ADD_DETOURS = "INSERT INTO `detours` (id, name, description, latitude, longitude) VALUES ( ?, ?, ?, ?, ?)"
	QUERY_FORMAT_ADD_DETOURS_EMOTION = "INSERT INTO `emotions` (emotions.id, emotions.happiness, emotions.natural, emotions.sadness, emotions.anger) VALUES ( ?, ?, ?, ?, ?)"
	QUERY_FORMAT_ADD_DETOURS_IMAGE = "INSERT INTO `images` (id, image) VALUES ( ?, ?)"
)

type GetSpotRequest struct {
	Latitude	float64	`json:"latitude"`
	Longitude	float64	`json:"longitude"`
	// spot_required
	// スポットへの所要時間
	Walktime	int		`json:"walktime"`
	Emotion		int		`json:"emotion"`
}

type GetSpotResponse struct {
	Spot	Spot	`json:"spot"`
}

type SpotInfo struct {
	ID			string
	Name		string
	Latitude	float64
	Longitude	float64
	// emotion table
	Happiness	float64
	Natural		float64
	Sadness		float64
	Anger		float64
}

type Spot struct {
	ID			string	`json:"id"`
	// 保留カラム
	Name		string	`json:"name"`
	Latitude	float64	`json:"latitude"`
	Longitude	float64	`json:"longitude"`
}

type GetDetourRequest struct {
	SpotLatitude	float64	`json:"spot_latitude"`
	SpotLongitude	float64	`json:"spot_longitude"`
	UserLatitude	float64	`json:"user_latitude"`
	UserLongitude	float64	`json:"user_longitude"`
	// spot_required
	// スポットへの所要時間
	Walktime	int		`json:"walktime"`
	Emotion		int		`json:"emotion"`
}

type GetDetourResponse struct {
	Detour	[]Detour	`json:"detour"`
}

type Detour struct {
	Name		string	`json:"name"`
	Image		string	`json:"image"`
	Description	string	`json:"description"`
	Latitude	float64	`json:"latitude"`
	Longitude	float64	`json:"longitude"`
}

type AddSpotRequest struct {
	Name		string	`json:"name"`
	Description	string	`json:"description"`
	Image		string	`json:"image"`
	Latitude	float64	`json:"latitude"`
	Longitude	float64	`json:"longitude"`
}

type Error struct {
	Code		int		`json:"code"`
	Message		string	`json:"message"`
	Description	string	`json:"description"`
}

type MutationResponse struct {
	Status	bool	`json:"status"`
}
