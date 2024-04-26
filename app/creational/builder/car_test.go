package builder

import "testing"

func TestCarBuilder(t *testing.T) {

	cb := &CarBuilder{}
	car := cb.Wheels().Seats().Structure().Build()

	if car.Seats != 4 || car.Wheels != 4 || car.Structure != "car" {
		t.Fatal("test case failed")
	}
}
