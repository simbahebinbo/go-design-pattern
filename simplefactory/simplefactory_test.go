package simplefactory

import "testing"

func TestExampleMouseFactory_CreateMouse(t *testing.T) {
	var mf MouseFactory
	m1 := mf.CreateMouse(0) //戴尔鼠标
	m1.SayMouseBrand()
	m2 := mf.CreateMouse(1) //惠普鼠标
	m2.SayMouseBrand()
	//OutPut:
	//Dell Mouse
	//Hp Mouse
}
