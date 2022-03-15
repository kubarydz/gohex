package api

import "github.com/kubarydz/go-hex/internal/ports"

type Adapter struct {
	db    ports.DbPort
	arith ports.ArithmeticPort
}

func NewAdapter(db ports.DbPort, arith ports.ArithmeticPort) *Adapter {
	return &Adapter{db, arith}
}

func (apia Adapter) GetAddition(a, b int32) (int32, error) {
	answ, err := apia.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddToHistory(answ, "addition")
	if err != nil {
		return 0, err
	}
	return answ, nil
}

func (apia Adapter) GetSubtraction(a, b int32) (int32, error) {
	answ, err := apia.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddToHistory(answ, "subtraction")
	if err != nil {
		return 0, err
	}

	return answ, nil
}

func (apia Adapter) GetMultiplication(a, b int32) (int32, error) {
	answ, err := apia.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddToHistory(answ, "multiplication")
	if err != nil {
		return 0, err
	}

	return answ, nil
}

func (apia Adapter) GetDivision(a, b int32) (int32, error) {
	answ, err := apia.arith.Division(a, b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddToHistory(answ, "division")
	if err != nil {
		return 0, err
	}

	return answ, nil
}
