package abstractfactory

type FamilyCar struct{}

func (f *FamilyCar) NumWheels() int { return 4 }
func (f *FamilyCar) NumSeats() int  { return 4 }
func (f *FamilyCar) NumDoors() int  { return 4 }
