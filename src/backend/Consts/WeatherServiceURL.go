package Consts

var weatherServiceURL string

func SetWeatherServiceURL(url string) {
	weatherServiceURL = url
}

func GetWeatherServiceURL() string {
	return weatherServiceURL
}
