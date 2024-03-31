package dtos

import "testing"

func TestWeatherDto_String(t *testing.T) {
	type fields struct {
		Location    string
		Temperature string
		Humidity    string
		SunHours    int
		WindSpeed   string
		Weather     string
		Date        string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Berlin String()",
			fields: fields{
				Location:    "Berlin",
				Temperature: "20°C",
				Humidity:    "50%",
				SunHours:    8,
				WindSpeed:   "5m/s",
				Weather:     "Sunny",
				Date:        "2021-09-01",
			},
			want: "Location: Berlin\nTemperature: 20°C\nHumidity: 50%\nSunHours: 8\nWindSpeed: 5m/s\nWeather: Sunny\nDate: 2021-09-01\n",
		},
		{
			name: "Test Paris String()",
			fields: fields{
				Location:    "Paris",
				Temperature: "25°C",
				Humidity:    "60%",
				SunHours:    10,
				WindSpeed:   "3m/s",
				Weather:     "Cloudy",
				Date:        "2021-09-01",
			},
			want: "Location: Paris\nTemperature: 25°C\nHumidity: 60%\nSunHours: 10\nWindSpeed: 3m/s\nWeather: Cloudy\nDate: 2021-09-01\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wd := WeatherDto{
				Location:    tt.fields.Location,
				Temperature: tt.fields.Temperature,
				Humidity:    tt.fields.Humidity,
				SunHours:    tt.fields.SunHours,
				WindSpeed:   tt.fields.WindSpeed,
				Weather:     tt.fields.Weather,
				Date:        tt.fields.Date,
			}
			if got := wd.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
