package coffe

import "fmt"

type Mocha struct {
	Coffe Coffe
}

func (m Mocha) GetCost() float64 {
	return m.Coffe.GetCost() + 1.00
}

func (m Mocha) GetDescription() string {
	return fmt.Sprintf("%v, Mocha", m.Coffe.GetDescription())
}

type Whipcream struct {
	Coffe Coffe
}

func (w Whipcream) GetCost() float64 {
	return w.Coffe.GetCost() + 0.1
}

func (w Whipcream) GetDescription() string {
	return fmt.Sprintf("%v, Whipcream", w.Coffe.GetDescription())
}
