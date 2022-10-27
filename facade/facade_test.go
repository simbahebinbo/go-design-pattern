package facade

import "testing"

func TestExampleComputer_StartUp(t *testing.T) {
	c := NewComputer()
	c.StartUp()
	//OutPut:
	//cpu startup
	//memory startup
}

func TestExampleComputer_ShutDown(t *testing.T) {
	c := NewComputer()
	c.ShutDown()
	//OutPut:
	//cpu shutdown
	//memory shutdown
}
