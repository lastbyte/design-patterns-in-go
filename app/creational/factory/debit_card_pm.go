package factory

type DebitCardPM struct{}

func (dpm *DebitCardPM) Pay(amount float64) string {
	if amount <= 0 {
		return "fail"
	}
	return "success"
}
