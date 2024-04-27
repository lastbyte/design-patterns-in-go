package abstractfactory

import "testing"

func TestMotorbikeFactory(t *testing.T) {
	fct := MotorbikeFactory{}

	_, err := fct.NewVehicle(1)

	if err != nil {
		t.Error("error in bike factory")
	}

	_, err = fct.NewVehicle(2)

	if err != nil {
		t.Error("error in bike factory")
	}

	_, err = fct.NewVehicle(3)

	if err == nil {
		t.Error("non existant bike should be nil")
	}
}
