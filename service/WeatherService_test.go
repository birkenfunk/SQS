package service

import (
	"reflect"
	"testing"

	"codeberg.org/Birkenfunk/SQS/consts"
	"codeberg.org/Birkenfunk/SQS/dtos"
)

func TestWeatherService_GetHealth(t *testing.T) {
	tests := []struct {
		name      string
		serverURL string
		wantErr   bool
	}{
		{
			name:      "Test GetHealth Success",
			serverURL: "http://external-api:3000",
			wantErr:   false,
		},
		{
			name:      "Test GetHealth Fail",
			serverURL: "http://external-api:3001",
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			consts.SetWeatherServiceURL(tt.serverURL)
			ws := &WeatherService{}
			if err := ws.GetHealth(); (err != nil) != tt.wantErr {
				t.Errorf("GetHealth() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWeatherService_GetWeather(t *testing.T) {
	type args struct {
		location string
	}
	tests := []struct {
		name      string
		serverURL string
		args      args
		want      *dtos.WeatherDto
		wantErr   bool
	}{
		{
			name:      "Test GetWeather Success",
			serverURL: "http://external-api:3000",
			args:      args{location: "test"},
			want: &dtos.WeatherDto{
				Location:    "Test",
				Temperature: "20°C",
				Humidity:    "50%",
				SunHours:    5,
				WindSpeed:   "10km/h",
				Weather:     "Sunny",
				Date:        "2024-01-01",
			},
			wantErr: false,
		},
		{
			name:      "Test GetWeather Fail",
			serverURL: "http://external-api:3001",
			args:      args{location: "test"},
			want:      nil,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			consts.SetWeatherServiceURL(tt.serverURL)
			ws := &WeatherService{}
			got, err := ws.GetWeather(tt.args.location)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWeather() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWeather() got = %v, want %v", got, tt.want)
			}
		})
	}
}
