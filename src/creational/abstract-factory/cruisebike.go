package abstractfactory

type CruiseBike struct{}

func (c *CruiseBike) NumWheels() int {
	return 2
}
func (c *CruiseBike) NumSeats() int {
	return 2
}
func (c *CruiseBike) MotorbikeType() int {
	return 2
}
