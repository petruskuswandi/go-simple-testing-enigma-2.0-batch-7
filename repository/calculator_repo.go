package repository

type CalculatorRepository interface {
	Add() (result *float64, err error)
	Sub() (result *float64, err error)
}

type calculatorRepository struct {
	Num1 float64
	Num2 float64
}

func (c *calculatorRepository) Add() (result *float64, err error) {
	subResult := c.Num1 + c.Num2
	result = &subResult
	err = nil
	return
}
func (c *calculatorRepository) Sub() (result *float64, err error) {
	subResult := c.Num1 - c.Num2
	result = &subResult
	err = nil
	return
}

func NewCalculatorRepository() CalculatorRepository {
	return &calculatorRepository{}
}
