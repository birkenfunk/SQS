package consts

import "testing"

func TestGetPort(t *testing.T) {
	tests := []struct {
		name string
		set  string
		want string
	}{
		{
			name: "Test GetPort",
			set:  "8080",
			want: "8080",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			port = tt.set
			if got := GetPort(); got != tt.want {
				t.Errorf("GetPort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWeatherServiceURL(t *testing.T) {
	tests := []struct {
		name string
		set  string
		want string
	}{
		{
			name: "Test GetWeatherServiceURL",
			set:  "http://localhost:8080",
			want: "http://localhost:8080",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			weatherServiceURL = tt.set
			if got := GetWeatherServiceURL(); got != tt.want {
				t.Errorf("GetWeatherServiceURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetPortFromString(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test SetPortFromString",
			args: args{
				p: "8080",
			},
			want: "8080",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetPortFromString(tt.args.p)
			if port != tt.want {
				t.Errorf("SetPortFromString() has set Port to %v, but should be %v", port, tt.want)
			}
		})
	}
}

func TestSetWeatherServiceURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test SetWeatherServiceURL",
			args: args{
				url: "http://localhost:8080",
			},
			want: "http://localhost:8080",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetWeatherServiceURL(tt.args.url)
			if weatherServiceURL != tt.want {
				t.Errorf("SetWeatherServiceURL() has set WeatherServiceURL to %v, but should be %v", weatherServiceURL, tt.want)
			}
		})
	}
}
