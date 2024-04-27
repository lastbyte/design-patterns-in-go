package abstractfactory

import "errors"

type CarFactory struct{}

func (cf *CarFactory) NewVehicle(v int) (Vehicle, error) {

	switch v {
	case 1:
		return new(LuxuryCar), nil
	case 2:
		return new(FamilyCar), nil
	default:
		return nil, errors.New("unknown_car")
	}
}
