package factory

type CashPM struct{}

func (cpm *CashPM) Pay(amount float64) string {
	if amount <= 0 {
		return "fail"
	}
	return "success"

}
