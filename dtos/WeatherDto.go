package dtos

import "strconv"

type WeatherDto struct {
	Location    string `json:"location"`
	Temperature string `json:"temperature"`
	Humidity    string `json:"humidity"`
	SunHours    int    `json:"sunHours"`
	WindSpeed   string `json:"windSpeed"`
	Weather     string `json:"weather"`
	Date        string `json:"date"`
}

// String
func (wd WeatherDto) String() string {
	return "Location: " + wd.Location + "\n" +
		"Temperature: " + wd.Temperature + "\n" +
		"Humidity: " + wd.Humidity + "\n" +
		"SunHours: " + strconv.Itoa(wd.SunHours) + "\n" +
		"WindSpeed: " + wd.WindSpeed + "\n" +
		"Weather: " + wd.Weather + "\n" +
		"Date: " + wd.Date + "\n"
}
