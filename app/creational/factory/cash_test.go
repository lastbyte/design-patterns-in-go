package factory

import "testing"

func TestPayCashMethod(t *testing.T) {

	cpm := GetPaymentMethod(1)
	msg := cpm.Pay(20.00)

	if msg == "" {
		t.Fatal("Error in making payment")
	}
	msg = cpm.Pay(-20.00)
	if msg != "fail" {
		t.Error("negative amount cannot be processed")
	}

}
