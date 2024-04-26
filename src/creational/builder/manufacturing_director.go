package builder

type ManufacturingDirector struct {
	Builder Builder
}

func (md *ManufacturingDirector) Construct() {
	md.Builder.Seats().Wheels().Structure()
}

func (md *ManufacturingDirector) SetBuilder(b Builder) {
	md.Builder = b
}
