package consts

var weatherServiceURL string

func SetWeatherServiceURL(url string) {
	weatherServiceURL = url
}

func GetWeatherServiceURL() string {
	return weatherServiceURL
}

var port string

func GetPort() string {
	return port
}

func SetPortFromString(p string) {
	port = p
}

var dbURL string

func SetDBURL(url string) {
	dbURL = url
}

func GetDBURL() string {
	return dbURL
}
