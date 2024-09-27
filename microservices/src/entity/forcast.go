package entity

type Forecast struct {
	DailyImpressions int `json:"daily_impressions"`
	DailyReach       int `json:"daily_reach"`
	Probability     float64 `json:"prediction"` 
}

type ForecastData struct {
	Forecast Forecast `json:"forecast"`
}