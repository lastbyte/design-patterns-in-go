package abstractfactory

import "testing"

func TestSportsBike(t *testing.T) {

	bike := SportsBike{}

	if bike.NumWheels() != 2 {
		t.Error("error in bike test function")
	} else if bike.NumSeats() != 2 {
		t.Error("error in bike test function")
	} else if bike.MototBikeType() != 1 {
		t.Error("error in bike test function")
	}
}
