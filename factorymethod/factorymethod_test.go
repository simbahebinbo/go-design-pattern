package factorymethod

import "testing"

func TestExampleDellMouseFactory_Create(t *testing.T) {
	var dmf DellMouseFactory
	dm := dmf.Create()
	dm.SayMouseBrand()
	//OutPut:
	//Dell Mouse
}

func TestExampleHpMouseFactory_Create(t *testing.T) {
	var hmf HpMouseFactory
	hm := hmf.Create()
	hm.SayMouseBrand()
	//OutPut:
	//Hp Mouse
}
