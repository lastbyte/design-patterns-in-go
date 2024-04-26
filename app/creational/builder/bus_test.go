package builder

import "testing"

func TestBusBuilder(t *testing.T) {

	bb := &BusBuilder{}
	bus := bb.Wheels().Seats().Structure().Build()

	if bus.Seats != 40 || bus.Wheels != 16 || bus.Structure != "Bus" {
		t.Fatal("test case failed")
	}
}
