package builder

import "testing"

func TestHouseBuilder_Build(t *testing.T) {
	house := &House{}
	houseBuilder := &Builder{house: house}
	h := houseBuilder.Color("red").Area(60.45).Material("wood").Build()
	t.Logf("%+v\n", h)
}
