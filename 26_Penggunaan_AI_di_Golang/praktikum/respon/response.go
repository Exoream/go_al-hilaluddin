package respon

type RecommendResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

type RecommendRequest struct {
	Budget     float64 `json:"budget"`
	Purpose    string  `json:"purpose"`
	Brand      string  `json:"brand"`
	RAM        string  `json:"ram"`
	CPU        string  `json:"cpu"`
	ScreenSize string  `json:"screen"`
}
