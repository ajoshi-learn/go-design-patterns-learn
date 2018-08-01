package builder

type BikeBuilder struct {
	v Vehicle
}

func (c *BikeBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 2
	return c
}

func (c *BikeBuilder) SetSeats() BuildProcess {
	c.v.Seats = 1
	return c
}

func (c *BikeBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Bike"
	return c
}

func (c *BikeBuilder) Build() Vehicle {
	return c.v
}
