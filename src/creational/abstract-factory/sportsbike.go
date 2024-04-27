package abstractfactory

type SportsBike struct{}

func (s *SportsBike) NumWheels() int {
	return 2
}
func (s *SportsBike) NumSeats() int {
	return 2
}
func (s *SportsBike) MototBikeType() int {
	return 1
}
