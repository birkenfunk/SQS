package logic

import (
	"fmt"
	"reflect"
	"testing"

	"codeberg.org/Birkenfunk/SQS/dtos"
	"codeberg.org/Birkenfunk/SQS/service"
)

var err = fmt.Errorf("failed to get weather")

func TestWeather_GetWeather(t *testing.T) {
	type args struct {
		location string
	}
	tests := []struct {
		name string
		args args
		want *dtos.WeatherDto
		mock *service.WeatherServiceMock
	}{
		{
			name: "Test GetWeather Success",
			args: args{
				location: "Berlin",
			},
			want: &dtos.WeatherDto{
				Location:    "Berlin",
				Temperature: "20°C",
				Humidity:    "20%",
				SunHours:    5,
				WindSpeed:   "50m/s",
				Weather:     "Sunny",
				Date:        "2021-09-01",
			},
			mock: &service.WeatherServiceMock{
				Weather: &dtos.WeatherDto{
					Location:    "Berlin",
					Temperature: "20°C",
					Humidity:    "20%",
					SunHours:    5,
					WindSpeed:   "50m/s",
					Weather:     "Sunny",
					Date:        "2021-09-01",
				},
				Error: nil,
			},
		},
		{
			name: "Test GetWeather Fail",
			args: args{
				location: "Berlin",
			},
			want: nil,
			mock: &service.WeatherServiceMock{
				Weather: nil,
				Error:   err,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Weather{}
			weatherService = tt.mock
			if got := w.GetWeather(tt.args.location); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWeather() = %v, want %v", got, tt.want)
			}
		})
	}
}
