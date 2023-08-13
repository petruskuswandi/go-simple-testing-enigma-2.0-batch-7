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
	calculator Calculator
	calMock    *CalculatorMock
}

func (suite *CalculatorTestSuite) SetupTest() {
	suite.calMock = new(CalculatorMock)
	suite.calculator = Calculator{
		Num1: 7,
		Num2: 1,
	}
}

func (suite *CalculatorTestSuite) TestCalculatorAdd_Success() {
	expected := 8.0
	suite.calMock.On("Add").Return(&expected, nil)
	actual, err := suite.calculator.Add()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, *actual)
}

func (suite *CalculatorTestSuite) TestCalculatorAdd_Fail() {
	expected := 10.0
	suite.calMock.On("Add").Return(&expected, nil)
	_, err := suite.calculator.Add()
	assert.Nil(suite.T(), err)
}

func (suite *CalculatorTestSuite) TestCalculatorSub_Success() {
	expected := 6.0
	suite.calMock.On("Sub").Return(&expected, nil)
	actual, err := suite.calculator.Sub()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, *actual)
}

func (suite *CalculatorTestSuite) TestCalculatorSub_Fail() {
	expected := 10.0
	suite.calMock.On("Sub").Return(&expected, nil)
	_, err := suite.calculator.Sub()
	assert.Nil(suite.T(), err)
}

func TestCalculatorSuite(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}
