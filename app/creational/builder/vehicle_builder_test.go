package builder

import "testing"

func TestCarBuilder(t *testing.T) {

	cb := &CarBuilder{}
	car := cb.Wheels().Seats().Structure().Build()

	if car.Seats != 4 || car.Wheels != 4 || car.Structure != "car" {
		t.Fatal("test case failed")
	}
}

func TestBusBuilder(t *testing.T) {

	bb := &BusBuilder{}
	bus := bb.Wheels().Seats().Structure().Build()

	if bus.Seats != 40 || bus.Wheels != 16 || bus.Structure != "Bus" {
		t.Fatal("test case failed")
	}
}
