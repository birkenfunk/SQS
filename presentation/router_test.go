package presentation

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"codeberg.org/Birkenfunk/SQS/business/handler"
	"codeberg.org/Birkenfunk/SQS/business/logic"
	"codeberg.org/Birkenfunk/SQS/dtos"
	"github.com/go-chi/chi/v5"
	"github.com/google/go-cmp/cmp"
)

// executeRequest, creates a new ResponseRecorder
// then executes the request by calling ServeHTTP in the router
// after which the handler writes the response to the response recorder
// which we can then inspect.
func executeRequest(req *http.Request, r *chi.Mux) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	return rr
}

func TestHealthEndpoint(t *testing.T) {
	r := InitRouter()
	req, err := http.NewRequest("GET", "/api/v1/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := executeRequest(req, r)
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}

	expected := "OK"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestWeatherEndpoint_Success(t *testing.T) {
	weatherDto := &dtos.WeatherDto{
		Location:    "Berlin",
		Temperature: "20°C",
		Humidity:    "20%",
		SunHours:    5,
		WindSpeed:   "50m/s",
		Weather:     "Sunny",
		Date:        "2021-09-01",
	}
	weatherMock := &logic.WeatherMock{
		WeatherDto: weatherDto,
	}
	handler.SetWeather(weatherMock)
	r := InitRouter()
	req, err := http.NewRequest("GET", "/api/v1/weather/berlin", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := executeRequest(req, r)
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
	if rr.Body.String() == "" {
		t.Errorf("handler returned empty body")
	}
	if rr.Header().Get("Content-Type") != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v", rr.Header().Get("Content-Type"), "application/json")
	}
	responseDto := &dtos.WeatherDto{}
	err = json.Unmarshal(rr.Body.Bytes(), responseDto)
	if err != nil {
		t.Errorf("failed to unmarshal response body")
	}
	if !cmp.Equal(responseDto, weatherDto) {
		t.Errorf("handler returned unexpected body: got %v want %v", responseDto, weatherDto)
	}
}

func TestWeatherEndpoint_Fail(t *testing.T) {
	weatherMock := &logic.WeatherMock{
		WeatherDto: nil,
	}
	handler.SetWeather(weatherMock)
	r := InitRouter()
	req, err := http.NewRequest("GET", "/api/v1/weather/berlin", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := executeRequest(req, r)
	if rr.Code != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusInternalServerError)
	}
	if rr.Body.String() != "Failed to get weather" {
		t.Errorf("handler returned wrong body: got %v want %v", rr.Body.String(), "Failed to get weather")
	}
}
