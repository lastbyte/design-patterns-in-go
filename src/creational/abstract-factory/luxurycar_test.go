package abstractfactory

import "testing"

func TestLuxuryCar(t *testing.T) {

	car := LuxuryCar{}

	if car.NumWheels() != 4 {
		t.Error("Error in luxury car ")
	}
	if car.NumSeats() != 4 {
		t.Error("Error in luxury car ")
	}
	if car.NumDoors() != 4 {
		t.Error("Error in luxury car ")
	}
}
