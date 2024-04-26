package factory

import "testing"

func TestPayDebinCardMethod(t *testing.T) {

	dpm := GetPaymentMethod(2)
	msg := dpm.Pay(20.00)
	if msg == "" {
		t.Error("Error in making payment")
	}
	msg = dpm.Pay(-20.00)
	if msg != "fail" {
		t.Error("negative amount cannot be processed")
	}

}
