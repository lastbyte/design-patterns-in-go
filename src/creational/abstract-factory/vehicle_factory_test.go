package abstractfactory

import "testing"

func TestNewVehicleFactory(t *testing.T) {

	carFactory, _ := NewVehicleFactory(1)

	if carFactory == nil {
		t.Error("car factory creation failed")
	}

	motorbikeFactory, _ := NewVehicleFactory(2)

	if motorbikeFactory == nil {
		t.Error("motorbike factory creation failed")
	}

	nonExistantFactory, _ := NewVehicleFactory(3)

	if nonExistantFactory != nil {
		t.Error("nonExistantFactory should be nil")
	}
}
