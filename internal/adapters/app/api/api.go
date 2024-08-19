package api

import "go-arch/internal/ports"

type Adapter struct {
	arith ports.ArithmeticPort
	db    ports.DbPort
}

func NewAdapter(db ports.DbPort, arith ports.ArithmeticPort) *Adapter {
	return &Adapter{arith: arith}
}

func (apia Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := apia.arith.Addition(a, b)
	err = apia.db.AddToHistory(answer, "addition")
	return answer, err
}

func (apia Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := apia.arith.Subtraction(a, b)
	err = apia.db.AddToHistory(answer, "subtraction")

	return answer, err
}

func (apia Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := apia.arith.Multiplication(a, b)
	err = apia.db.AddToHistory(answer, "multiplication")

	return answer, err
}

func (apia Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := apia.arith.Division(a, b)
	err = apia.db.AddToHistory(answer, "division")

	return answer, err
}
