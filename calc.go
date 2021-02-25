package awesomeProject24

type CalcService struct{}


func (s *CalcService) Add(a, b int) int {
	return a + b
}

func (s *CalcService) Subtract(a, b int) int {
	return a - b
}
