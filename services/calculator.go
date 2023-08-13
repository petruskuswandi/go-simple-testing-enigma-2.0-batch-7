package services

type Calculator struct {
	Num1 float64
	Num2 float64
}

func (c *Calculator) Add() (result *float64, err error) {
	subResult := c.Num1 + c.Num2
	result = &subResult
	err = nil
	return
}
func (c *Calculator) Sub() (result *float64, err error) {
	subResult := c.Num1 - c.Num2
	result = &subResult
	err = nil
	return
}
