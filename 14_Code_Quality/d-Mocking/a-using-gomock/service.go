package main

type Service struct {
	calc Calculator
}

func NewService(calc Calculator) *Service {
	return &Service{calc: calc}
}

func (s *Service) PerformAddition(a, b int) int {
	return s.calc.Add(a, b)
}


// Generate the Mock 
// mockgen -source=calculator.go -destination=mock_calculator.go -package=main
