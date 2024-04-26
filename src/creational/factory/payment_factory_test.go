package factory

import "testing"

func TestGetPaymentMethod(t *testing.T) {

	cpm := GetPaymentMethod(1)

	if cpm == nil {
		t.Error("error in creating cash payment method")
	}

	dpm := GetPaymentMethod(2)

	if dpm == nil {
		t.Error("error in creating debit card cash method")
	}

	nonExisting := GetPaymentMethod(3)

	if nonExisting != nil {
		t.Error("Non existing payment method sould be nil")
	}

}
