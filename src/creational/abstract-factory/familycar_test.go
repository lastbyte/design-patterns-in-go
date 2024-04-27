package abstractfactory

import "testing"

func TestFamilyCar(t *testing.T) {
	car := FamilyCar{}

	if car.NumWheels() != 4 {
		t.Error("Family car error")
	} else if car.NumSeats() != 4 {
		t.Error("Family car error")
	} else if car.NumDoors() != 4 {
		t.Error("Family car error")
	}
}
