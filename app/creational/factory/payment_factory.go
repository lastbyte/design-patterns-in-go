package factory

type PaymentMethodFactory interface {
	GetPaymentMethod(paymentType int) PaymentMethod
}

const (
	CASH       = 1
	DEBIT_CARD = 2
)

type PaymentMethod interface {
	Pay(amount float64) string
}

func GetPaymentMethod(paymentType int) PaymentMethod {
	switch paymentType {
	case CASH:
		cpm := &CashPM{}
		return cpm
	case DEBIT_CARD:
		dpm := &DebitCardPM{}
		return dpm
	}
	return nil
}
