package builder

type Builder interface {
	Wheels() Builder
	Seats() Builder
	Structure() Builder
	Build() Vehicle
}

type BusBuilder struct {
	Vehicle Vehicle
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
