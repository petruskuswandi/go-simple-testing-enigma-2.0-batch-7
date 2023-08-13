package services

import "go-simple-testing/repository"

type CalculatorService interface {
	CalculatorAdd() (result *float64, err error)
	CalculatorSub() (result *float64, err error)
}

type calculatorService struct {
	repo repository.CalculatorRepository
}

// CalculatorAdd implements CalculatorService
func (c *calculatorService) CalculatorAdd() (result *float64, err error) {
	return c.repo.Add()
}

// CalculatorSub implements CalculatorService
func (c *calculatorService) CalculatorSub() (result *float64, err error) {
	return c.repo.Sub()
}

func NewCalculatorService(repo repository.CalculatorRepository) CalculatorService {
	return &calculatorService{repo: repo}
}
