package abstractfactory

import "testing"

func TestCarFactory(t *testing.T) {

	cft := &CarFactory{}

	luxuryCar, _ := cft.NewVehicle(1)

	if luxuryCar == nil {
		t.Error("error in car factory")
	}

	familyCar, _ := cft.NewVehicle(2)

	if familyCar == nil {
		t.Error("error in car factory")
	}

	nonExistantCar, _ := cft.NewVehicle(3)

	if nonExistantCar != nil {
		t.Error("non existing car should be nil")
	}

}
