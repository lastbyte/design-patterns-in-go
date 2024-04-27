package abstractfactory

import "testing"

func TestCruiseBike(t *testing.T) {

	bike := CruiseBike{}

	if bike.NumSeats() != 2 {
		t.Error("error in cruise bike")
	} else if bike.NumWheels() != 2 {
		t.Error("error in cruise bike")
	} else if bike.MotorbikeType() != 2 {
		t.Error("error in cruise bike")
	}
}
