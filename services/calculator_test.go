package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CalculatorMock struct {
	mock.Mock
}

func (c *CalculatorMock) Add() (result *float64, err error) {
	args := c.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*float64), nil
}

func (c *CalculatorMock) Sub() (result *float64, err error) {
	args := c.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*float64), nil
}

type CalculatorTestSuite struct {
	suite.Suite
	calMock *CalculatorMock
}

func (suite *CalculatorTestSuite) SetupTest() {
	suite.calMock = new(CalculatorMock)
}

func (suite *CalculatorTestSuite) TestCalculatorAdd_Success() {
	expected := 8.0
	suite.calMock.On("Add").Return(&expected, nil)
	calSrv := NewCalculatorService(suite.calMock)
	actual, err := calSrv.CalculatorAdd()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, *actual)
}

func (suite *CalculatorTestSuite) TestCalculatorAdd_Fail() {
	expected := 10.0
	suite.calMock.On("Add").Return(&expected, nil)
	calSrv := NewCalculatorService(suite.calMock)
	_, err := calSrv.CalculatorAdd()
	assert.Nil(suite.T(), err)
}

func (suite *CalculatorTestSuite) TestCalculatorSub_Success() {
	expected := 6.0
	suite.calMock.On("Sub").Return(&expected, nil)
	calSrv := NewCalculatorService(suite.calMock)
	actual, err := calSrv.CalculatorSub()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, *actual)
}

func (suite *CalculatorTestSuite) TestCalculatorSub_Fail() {
	expected := 10.0
	suite.calMock.On("Sub").Return(&expected, nil)
	calSrv := NewCalculatorService(suite.calMock)
	_, err := calSrv.CalculatorSub()
	assert.Nil(suite.T(), err)
}

func TestCalculatorTestSuite(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}
