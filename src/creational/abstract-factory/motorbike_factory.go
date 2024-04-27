package abstractfactory

import "errors"

type MotorbikeFactory struct{}

func (mf *MotorbikeFactory) NewVehicle(v int) (Motorbike, error) {
	switch v {
	case 1:
		return new(SportsBike), nil
	case 2:
		return new(CruiseBike), nil
	default:
		return nil, errors.New("unknown_motorbike")
	}
}
