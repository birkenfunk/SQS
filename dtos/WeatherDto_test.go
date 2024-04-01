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

func TestWeatherDto_Equals(t *testing.T) {
	type fields struct {
		Location    string
		Temperature string
		Humidity    string
		SunHours    int
		WindSpeed   string
		Weather     string
		Date        string
	}
	type args struct {
		other *WeatherDto
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Test Berlin Equals()",
			fields: fields{
				Location:    "Berlin",
				Temperature: "20°C",
				Humidity:    "50%",
				SunHours:    8,
				WindSpeed:   "5m/s",
				Weather:     "Sunny",
				Date:        "2021-09-01",
			},
			args: args{
				other: &WeatherDto{
					Location:    "Berlin",
					Temperature: "20°C",
					Humidity:    "50%",
					SunHours:    8,
					WindSpeed:   "5m/s",
					Weather:     "Sunny",
					Date:        "2021-09-01",
				},
			},
			want: true,
		},
		{
			name: "Test Paris Equals()",
			fields: fields{
				Location:    "Paris",
				Temperature: "25°C",
				Humidity:    "60%",
				SunHours:    10,
				WindSpeed:   "3m/s",
				Weather:     "Cloudy",
				Date:        "2021-09-01",
			},
			args: args{
				other: &WeatherDto{
					Location:    "Paris",
					Temperature: "25°C",
					Humidity:    "60%",
					SunHours:    10,
					WindSpeed:   "3m/s",
					Weather:     "Cloudy",
					Date:        "2021-09-01",
				},
			},
			want: true,
		},
		{
			name: "Test Berlin and Paris Equals()",
			fields: fields{
				Location:    "Berlin",
				Temperature: "20°C",
				Humidity:    "50%",
				SunHours:    8,
				WindSpeed:   "5m/s",
				Weather:     "Sunny",
				Date:        "2021-09-01",
			},
			args: args{
				other: &WeatherDto{
					Location:    "Paris",
					Temperature: "25°C",
					Humidity:    "60%",
					SunHours:    10,
					WindSpeed:   "3m/s",
					Weather:     "Cloudy",
					Date:        "2021-09-01",
				},
			},
			want: false,
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
			if got := wd.Equals(tt.args.other); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeatherDto_Diff(t *testing.T) {
	type fields struct {
		Location    string
		Temperature string
		Humidity    string
		SunHours    int
		WindSpeed   string
		Weather     string
		Date        string
	}
	type args struct {
		other *WeatherDto
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Test Berlin Diff()",
			fields: fields{
				Location:    "Berlin",
				Temperature: "20°C",
				Humidity:    "50%",
				SunHours:    8,
				WindSpeed:   "5m/s",
				Weather:     "Sunny",
				Date:        "2021-09-01",
			},
			args: args{
				other: &WeatherDto{
					Location:    "Berlin",
					Temperature: "20°C",
					Humidity:    "50%",
					SunHours:    8,
					WindSpeed:   "5m/s",
					Weather:     "Sunny",
					Date:        "2021-09-01",
				},
			},
			want: "",
		},
		{
			name: "Test Paris Diff()",
			fields: fields{
				Location:    "Paris",
				Temperature: "25°C",
				Humidity:    "60%",
				SunHours:    10,
				WindSpeed:   "3m/s",
				Weather:     "Cloudy",
				Date:        "2021-09-01",
			},
			args: args{
				other: &WeatherDto{
					Location:    "Paris",
					Temperature: "25°C",
					Humidity:    "60%",
					SunHours:    10,
					WindSpeed:   "3m/s",
					Weather:     "Cloudy",
					Date:        "2021-09-01",
				},
			},
			want: "",
		},
		{
			name: "Test Berlin and Paris Diff()",
			fields: fields{
				Location:    "Berlin",
				Temperature: "20°C",
				Humidity:    "50%",
				SunHours:    8,
				WindSpeed:   "5m/s",
				Weather:     "Sunny",
				Date:        "2021-09-01",
			},
			args: args{
				other: &WeatherDto{
					Location:    "Paris",
					Temperature: "25°C",
					Humidity:    "60%",
					SunHours:    10,
					WindSpeed:   "3m/s",
					Weather:     "Cloudy",
					Date:        "2022-09-01",
				},
			},
			want: "Location: Berlin != Paris\nTemperature: 20°C != 25°C\nHumidity: 50% != 60%\nSunHours: 8 != 10\nWindSpeed: 5m/s != 3m/s\nWeather: Sunny != Cloudy\nDate: 2021-09-01 != 2022-09-01\n",
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
			if got := wd.Diff(tt.args.other); got != tt.want {
				t.Errorf("Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}
