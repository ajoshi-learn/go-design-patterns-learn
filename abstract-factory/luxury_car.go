package abstract_factory

type LuxuryCar struct {
}

func (c *LuxuryCar) NumDoors() int {
	return 4
}

func (c *LuxuryCar) NumWheels() int {
	return 4
}

func (c *LuxuryCar) NumSeats() int {
	return 5
}
