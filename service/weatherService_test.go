package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"testing/iotest"

	"codeberg.org/Birkenfunk/SQS/dtos"
	"codeberg.org/Birkenfunk/SQS/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type NonClosableBody struct {
	data  []byte
	index int
}

func (b *NonClosableBody) Close() error {
	return err
}

func (b *NonClosableBody) Read(_ []byte) (n int, err error) {
	if b.index >= len(b.data) {
		return 0, io.EOF
	}
	n = copy(b.data, b.data[b.index:])
	b.index += n
	return n, nil
}

type WeatherServiceSuite struct {
	suite.Suite
	mockClient *mocks.HTTPClient
	service    *WeatherService
	dto        *dtos.WeatherDto
}

func TestWeatherServiceSuite(t *testing.T) {
	suite.Run(t, new(WeatherServiceSuite))
}

func (suite *WeatherServiceSuite) SetupTest() {
	suite.mockClient = new(mocks.HTTPClient)
	suite.service = &WeatherService{
		client: suite.mockClient,
	}
	suite.dto = &dtos.WeatherDto{
		Location:    "Berlin",
		Temperature: "20°C",
		Humidity:    "50%",
		SunHours:    8,
		WindSpeed:   "5m/s",
		Weather:     "Sunny",
		Date:        "2021-09-01",
	}
}

var err = fmt.Errorf("failed to get weather")

func (suite *WeatherServiceSuite) TestGetWeather_Success() {
	// given:
	dto, err := json.Marshal(suite.dto)
	suite.Require().NoError(err)
	body := io.NopCloser(bytes.NewReader(dto))
	response := http.Response{Body: body, StatusCode: 200}
	suite.mockClient.On("Do", mock.Anything).Return(&response, nil)

	// when:
	result, err := suite.service.GetWeather("Berlin")

	// then:
	suite.Require().NoError(err)
	suite.Equal(suite.dto, result)
}

func (suite *WeatherServiceSuite) TestGetWeather_Fail() {
	// given:
	suite.mockClient.On("Do", mock.Anything).Return(nil, err)

	// when:
	result, err := suite.service.GetWeather("Berlin")

	// then:
	suite.Nil(result)
	suite.Error(err)
}

func (suite *WeatherServiceSuite) TestGetWeather_FailStatusCode() {
	// given:
	body := io.NopCloser(bytes.NewReader([]byte("error")))
	response := http.Response{Body: body, StatusCode: 500}
	suite.mockClient.On("Do", mock.Anything).Return(&response, nil)

	// when:
	result, err := suite.service.GetWeather("Berlin")

	// then:
	suite.Nil(result)
	suite.Error(err)
}

func (suite *WeatherServiceSuite) TestGetWeather_FailUnmarshal() {
	// given:
	body := io.NopCloser(bytes.NewReader([]byte("error")))
	response := http.Response{Body: body, StatusCode: 200}
	suite.mockClient.On("Do", mock.Anything).Return(&response, nil)

	// when:
	result, err := suite.service.GetWeather("Berlin")

	// then:
	suite.Nil(result)
	suite.Error(err)
}

func (suite *WeatherServiceSuite) TestGetWeather_FailClose() {
	// given:
	dto, err := json.Marshal(suite.dto)
	suite.Require().NoError(err)
	body := &NonClosableBody{data: dto}
	response := http.Response{Body: body, StatusCode: 200}
	suite.mockClient.On("Do", mock.Anything).Return(&response, nil)
	suite.mockClient.On("Do", mock.Anything).Return(nil, err)

	// when:
	result, err := suite.service.GetWeather("Berlin")

	// then:
	suite.Nil(result)
	suite.Error(err)
}

func (suite *WeatherServiceSuite) TestGetWeather_FailReadAll() {
	// given:
	body := io.NopCloser(iotest.ErrReader(err))
	response := http.Response{Body: body, StatusCode: 200}
	suite.mockClient.On("Do", mock.Anything).Return(&response, nil)
	suite.mockClient.On("Do", mock.Anything).Return(nil, err)

	// when:
	result, err := suite.service.GetWeather("Berlin")

	// then:
	suite.Nil(result)
	suite.Error(err)
}

func (suite *WeatherServiceSuite) TestGetHealth_Success() {
	// given:
	body := io.NopCloser(bytes.NewReader([]byte("OK")))
	response := http.Response{Body: body, StatusCode: 200}
	suite.mockClient.On("Do", mock.Anything).Return(&response, nil)

	// when:
	err := suite.service.GetHealth()

	// then:
	suite.NoError(err)
}

func (suite *WeatherServiceSuite) TestGetHealth_Fail() {
	// given:
	suite.mockClient.On("Do", mock.Anything).Return(nil, err)

	// when:
	err := suite.service.GetHealth()

	// then:
	suite.Error(err)
}

func (suite *WeatherServiceSuite) TestGetHealth_FailStatusCode() {
	// given:
	body := io.NopCloser(bytes.NewReader([]byte("error")))
	response := http.Response{Body: body, StatusCode: 500}
	suite.mockClient.On("Do", mock.Anything).Return(&response, nil)

	// when:
	err := suite.service.GetHealth()

	// then:
	suite.Error(err)
}

func (suite *WeatherServiceSuite) TestGetHealth_FailClose() {
	// given:
	body := &NonClosableBody{data: []byte("OK")}
	response := http.Response{Body: body, StatusCode: 200}
	suite.mockClient.On("Do", mock.Anything).Return(&response, nil)
	suite.mockClient.On("Do", mock.Anything).Return(nil, err)

	// when:
	err := suite.service.GetHealth()

	// then:
	suite.Error(err)
}
