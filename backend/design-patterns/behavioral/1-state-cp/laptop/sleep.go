package laptop

type Sleeping struct {
	Laptop *Laptop
}

func (s Sleeping) Press() {
	s.Laptop.ChangeCurrentState("On")
	s.Laptop.ChangeState(On{s.Laptop})
}

func (s Sleeping) CanTurnOnLaptop() bool {
	return true
}

func (s Sleeping) Sleep() {
	return
}
