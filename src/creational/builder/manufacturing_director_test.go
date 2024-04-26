package builder

import "testing"

func TestSetBuilder(t *testing.T) {

	md := ManufacturingDirector{}
	b := &CarBuilder{}

	md.SetBuilder(b)

	if md.Builder != b {
		t.Fatal("test failed in setting builder")
	}
}

func TestConstruct(t *testing.T) {
	md := ManufacturingDirector{}
	b := &CarBuilder{}

	md.SetBuilder(b)

	if md.Builder != b {
		t.Fatal("test failed in setting builder")
	}
	md.Construct()
	car := b.Build()

	if car.Seats != 4 || car.Wheels != 4 || car.Structure != "car" {
		t.Fatal("test failing in running construct")
	}

}
