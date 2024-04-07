package consts

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ConstantsSuite struct {
	suite.Suite
}

func TestConstantsSuite(t *testing.T) {
	suite.Run(t, &ConstantsSuite{})
}

func (suite *ConstantsSuite) SetupTest() {
	port = ""
	weatherServiceURL = ""
}

func (suite *ConstantsSuite) TearDownSuite() {
	port = ""
	weatherServiceURL = ""
}

func (suite *ConstantsSuite) TestGetPort() {
	// given:
	port = "8080"
	SetPortFromString(port)

	// when:
	result := GetPort()

	// then:
	suite.Equal(port, result)
}

func (suite *ConstantsSuite) TestGetWeatherServiceURL() {
	// given:
	weatherServiceURL = "http://localhost:8080"
	SetWeatherServiceURL(weatherServiceURL)

	// when:
	result := GetWeatherServiceURL()

	// then:
	suite.Equal(weatherServiceURL, result)
}

func (suite *ConstantsSuite) TestSetPortFromString() {
	// given:
	port = "8080"

	// when:
	SetPortFromString(port)

	// then:
	suite.Equal(port, GetPort())
}

func (suite *ConstantsSuite) TestSetWeatherServiceURL() {
	// given:
	weatherServiceURL = "http://localhost:8080"

	// when:
	SetWeatherServiceURL(weatherServiceURL)

	// then:
	suite.Equal(weatherServiceURL, GetWeatherServiceURL())
}

func (suite *ConstantsSuite) TestSetDBURL() {
	// given:
	dbURL := "localhost:6379"

	// when:
	SetDBURL(dbURL)

	// then:
	suite.Equal(dbURL, GetDBURL())
}

func (suite *ConstantsSuite) TestGetDBURL() {
	// given:
	dbURL := "localhost:6379"
	SetDBURL(dbURL)

	// when:
	result := GetDBURL()

	// then:
	suite.Equal(dbURL, result)
}
