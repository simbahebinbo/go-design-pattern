package abstractfactory

func ExampleDellFactory_CreateMouse() {
	df := new(DellFactory)
	df.CreateMouse()
	df.CreateKeyBoard()
	df.Mouse.SayMouseBrand()
	df.KeyBoard.SayKeyBoardBrand()
	//OutPut:
	//Hp Mouse
	//Dell KeyBoard
}
