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

func (wd WeatherDto) Equals(other *WeatherDto) bool {
	return wd.Location == other.Location &&
		wd.Temperature == other.Temperature &&
		wd.Humidity == other.Humidity &&
		wd.SunHours == other.SunHours &&
		wd.WindSpeed == other.WindSpeed &&
		wd.Weather == other.Weather &&
		wd.Date == other.Date
}

func (wd WeatherDto) Diff(other *WeatherDto) string {
	diff := ""
	if wd.Location != other.Location {
		diff += "Location: " + wd.Location + " != " + other.Location + "\n"
	}
	if wd.Temperature != other.Temperature {
		diff += "Temperature: " + wd.Temperature + " != " + other.Temperature + "\n"
	}
	if wd.Humidity != other.Humidity {
		diff += "Humidity: " + wd.Humidity + " != " + other.Humidity + "\n"
	}
	if wd.SunHours != other.SunHours {
		diff += "SunHours: " + strconv.Itoa(wd.SunHours) + " != " + strconv.Itoa(other.SunHours) + "\n"
	}
	if wd.WindSpeed != other.WindSpeed {
		diff += "WindSpeed: " + wd.WindSpeed + " != " + other.WindSpeed + "\n"
	}
	if wd.Weather != other.Weather {
		diff += "Weather: " + wd.Weather + " != " + other.Weather + "\n"
	}
	if wd.Date != other.Date {
		diff += "Date: " + wd.Date + " != " + other.Date + "\n"
	}
	return diff
}
