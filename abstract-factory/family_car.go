package abstract_factory

type FamilyCar struct {
}

func (c *FamilyCar) NumDoors() int {
	return 5
}

func (c *FamilyCar) NumWheels() int {
	return 4
}

func (c *FamilyCar) NumSeats() int {
	return 5
}
