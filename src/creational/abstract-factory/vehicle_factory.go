package abstractfactory

import "errors"

type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

func NewVehicleFactory(v int) (VehicleFactory, error) {
	switch v {
	case 1:
		return new(CarFactory), nil
	case 2:
		return new(MotorbikeFactory), nil
	default:
		return nil, errors.New("unknown_factory")
	}
}
