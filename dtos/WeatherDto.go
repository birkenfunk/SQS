package dtos

type WeatherDto struct {
	Location    string `json:"location"`
	Temperature string `json:"temperature"`
	Humidity    string `json:"humidity"`
	SunHours    int    `json:"sunHours"`
	WindSpeed   string `json:"windSpeed"`
	Weather     string `json:"weather"`
	Date        string `json:"date"`
}
