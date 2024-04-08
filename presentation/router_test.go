package presentation

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"codeberg.org/Birkenfunk/SQS/business/handler"
	"codeberg.org/Birkenfunk/SQS/dtos"
	"codeberg.org/Birkenfunk/SQS/mocks"
	"codeberg.org/Birkenfunk/SQS/testfixtures"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/suite"
)

type RouterSuite struct {
	suite.Suite
	router  IRouter
	weather *mocks.IWeather
}

func TestRouterSuite(t *testing.T) {
	suite.Run(t, &RouterSuite{})
}

func (rs *RouterSuite) SetupTest() {
	rs.weather = new(mocks.IWeather)
	weatherHandler := handler.WeatherHandler{}
	weatherHandler.SetWeather(rs.weather)
	rs.router = &Router{
		weatherHandler: &weatherHandler,
		healthHandler:  handler.NewHealthHandler(),
	}
}

// executeRequest, creates a new ResponseRecorder
// then executes the request by calling ServeHTTP in the router
// after which the handler writes the response to the response recorder
// which we can then inspect.
func executeRequest(req *http.Request, r *chi.Mux) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	return rr
}

func (rs *RouterSuite) TestNewRouter() {
	testfixtures.SetUpAllVariables()
	router := NewRouter()
	rs.Require().NotNil(router)
	rs.Require().IsType(&Router{}, router)
}

func (rs *RouterSuite) TestHealthEndpoint() {
	req, err := http.NewRequest("GET", "/api/v1/health", nil)
	rs.Require().NoError(err)

	rr := executeRequest(req, rs.router.InitRouter())
	rs.Require().Equal(http.StatusOK, rr.Code)
	rs.Require().Equal("OK", rr.Body.String())
}

func (rs *RouterSuite) TestWeatherEndpoint_Success() {
	// given:
	weatherDto := &dtos.WeatherDto{
		Location:    "Berlin",
		Temperature: "20°C",
		Humidity:    "20%",
		SunHours:    5,
		WindSpeed:   "50m/s",
		Weather:     "Sunny",
		Date:        "2021-09-01",
	}
	rs.weather.On("GetWeather", "berlin").Return(weatherDto)
	// when:
	req, err := http.NewRequest("GET", "/api/v1/weather/berlin", nil)
	rs.Require().NoError(err)
	rr := executeRequest(req, rs.router.InitRouter())

	// then:
	rs.Require().Equal(http.StatusOK, rr.Code)
	rs.Require().Equal("application/json", rr.Header().Get("Content-Type"))
	rs.Require().NotNil(rr.Body)

	responseDto := &dtos.WeatherDto{}
	err = json.Unmarshal(rr.Body.Bytes(), responseDto)
	rs.Require().NoError(err)
	rs.Require().Equal(weatherDto, responseDto)
}

func (rs *RouterSuite) TestWeatherEndpoint_Fail() {
	// given:
	rs.weather.On("GetWeather", "berlin").Return(nil)

	// when:
	req, err := http.NewRequest("GET", "/api/v1/weather/berlin", nil)
	rs.Require().NoError(err)

	rr := executeRequest(req, rs.router.InitRouter())

	// then:
	rs.Require().Equal(http.StatusInternalServerError, rr.Code)
	rs.Require().Equal("Failed to get weather", rr.Body.String())
}
