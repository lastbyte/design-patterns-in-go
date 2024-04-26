package builder

type CarBuilder struct {
	Vehicle Vehicle
}

func (c *CarBuilder) Wheels() Builder {
	c.Vehicle.Wheels = 4
	return c
}
func (c *CarBuilder) Seats() Builder {
	c.Vehicle.Seats = 4
	return c
}
func (c *CarBuilder) Structure() Builder {
	c.Vehicle.Structure = "car"
	return c
}

func (c *CarBuilder) Build() Vehicle {
	return c.Vehicle
}
