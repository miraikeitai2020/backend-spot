package model

/*--------------------------------*/
/*        SQL Query Format        */
/*--------------------------------*/
const(
	QUERY_FORMAT_GET_SPOT = "SELECT spots.id, spots.name, spots.latitude, spots.longitude, emotions.happiness, emotions.calm, emotions.exciting, emotions.sadness FROM `spots`, `emotions` WHERE spots.id = emotions.id GROUP BY spots.id"
	QUERY_FORMAT_GET_DETOURS = "SELECT detours.id, detours.name, detours.latitude, detours.longitude, emotions.happiness, emotions.calm, emotions.exciting, emotions.sadness FROM `detours`, `emotions` WHERE detours.id = emotions.id GROUP BY detours.id"
	QUERY_FORMAT_ADD_DETOURS = "INSERT INTO `detours` (id, name, description, latitude, longitude) VALUES (?, ?, ?, ?, ?)"
	QUERY_FORMAT_ADD_DETOURS_IMAGE = "INSERT INTO `images` (id, image) VALUES (?, ?)"
	QUERY_FORMAT_ADD_DETOURS_EMOTIONS="INSERT INTO `emotions` (id,emotions.happiness,emotions.calm,emotions.exciting,emotions.sadness) VALUES(?,?,?,?,?)"
)

/*-----------------------------------*/
/*     SQL Query Response Struct     */
/*-----------------------------------*/
type SpotInfo struct {
	// `spot` table data
	ID			string
	Name		string
	Latitude	float64
	Longitude	float64
	// `emotions` table data
	Happiness	float64
	Calm		float64
	Exciting 	float64
	Sadness		float64
}

type DetourInfo struct {
	// `spot` table data
	ID			string
	Name		string
	Latitude	float64
	Longitude	float64
	Image		string	
	// `emotions` table data
	Happiness	float64
	Calm		float64
	Exciting 	float64
	Sadness		float64
}
type Spot struct {
	ID			string	`json:"id"`
	Name		string	`json:"name"`
	Latitude	float64	`json:"latitude"`
	Longitude	float64	`json:"longitude"`
}

type Detour struct {
	ID			string	`json:"id"`
	Name		string	`json:"name"`
	Image		string	`json:"image"`
	Description	string	`json:"description"`
	Latitude	float64	`json:"latitude"`
	Longitude	float64	`json:"longitude"`
}

/*--------------------------------*/
/*     Handler Request Struct     */
/*--------------------------------*/
type GetSpotRequest struct {
	Latitude	float64	`json:"latitude"`
	Longitude	float64	`json:"longitude"`
	// スポットへの所要時間
	Walktime	int		`json:"walktime"`
	Emotion		int		`json:"emotion"`
}

type GetDetourRequest struct {
	SpotLatitude	float64	`json:"spot_latitude"`
	SpotLongitude	float64	`json:"spot_longitude"`
	UserLatitude	float64	`json:"user_latitude"`
	UserLongitude	float64	`json:"user_longitude"`
	// スポットへの所要時間
	Walktime	int		`json:"walktime"`
	Emotion		int		`json:"emotion"`
}

type AddSpotRequest struct {
	Name		string	`json:"name"`
	Image		string	`json:"image"`
	Description	string	`json:"description"`
	Latitude	float64	`json:"latitude"`
	Longitude	float64	`json:"longitude"`
	Emotion		int		`json:"emotion"`

}
