package perusahaan

type CTO struct {
	Subordinate []Employee
}

func (c CTO) GetSalary() int {
	return 30
}

func (c CTO) TotalDivisonSalary() int {
	total := c.GetSalary()
	for _, employee := range c.Subordinate {
		total += employee.TotalDivisonSalary()
	}
	return total
}
