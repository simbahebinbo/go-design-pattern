package abstractfactory

import "testing"

func TestExampleDellFactory_CreateMouse(t *testing.T) {
	df := new(DellFactory)
	df.CreateMouse()
	df.CreateKeyBoard()
	df.Mouse.SayMouseBrand()
	df.KeyBoard.SayKeyBoardBrand()
	//OutPut:
	//Hp Mouse
	//Dell KeyBoard
}
