package builder

type Vehicle struct {
	Wheels    int
	Seats     int
	Structure string
}

type Builder interface {
	Wheels() Builder
	Seats() Builder
	Structure() Builder
	Build() Vehicle
}

type CarBuilder struct {
	Vehicle Vehicle
}
type BusBuilder struct {
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

func (b *BusBuilder) Wheels() Builder {
	b.Vehicle.Wheels = 16
	return b
}
func (b *BusBuilder) Seats() Builder {
	b.Vehicle.Seats = 40
	return b
}
func (b *BusBuilder) Structure() Builder {
	b.Vehicle.Structure = "Bus"
	return b
}

func (b *BusBuilder) Build() Vehicle { return b.Vehicle }
