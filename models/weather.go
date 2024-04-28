package models

type WeatherAPI struct {
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	Timezone       string  `json:"timezone"`
	TimezoneOffset int64   `json:"timezone_offset"`
	Current        struct {
		Clouds     int64   `json:"clouds"`
		DewPoint   float64 `json:"dew_point"`
		Dt         int64   `json:"dt"`
		FeelsLike  float64 `json:"feels_like"`
		Humidity   int64   `json:"humidity"`
		Pressure   int64   `json:"pressure"`
		Sunrise    int64   `json:"sunrise"`
		Sunset     int64   `json:"sunset"`
		Temp       float64 `json:"temp"`
		Uvi        float64 `json:"uvi"`
		Visibility int64   `json:"visibility"`
		Weather    []struct {
			Description string `json:"description"`
			Icon        string `json:"icon"`
			ID          int64  `json:"id"`
			Main        string `json:"main"`
		} `json:"weather"`
		WindDeg   int64   `json:"wind_deg"`
		WindSpeed float64 `json:"wind_speed"`
	} `json:"current"`
}
