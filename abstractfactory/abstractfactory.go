// 抽象工厂模式
package abstractfactory

import "fmt"

//意图：提供一个创建一系列相关或相互依赖对象的接口，而无需指定它们具体的类,用于生成产品族的工厂，所生成的对象是有关联的
//解决:主要解决接口选择的问题

//不足:产品族扩展非常困难，要增加一个系列的某一产品，既要在抽象的 Creator 里加代码，又要在具体的里面加代码。

//实例:PC厂商是一个抽象工厂，是鼠标和键盘的组合，鼠标和键盘又是不同的工厂生产，抽象工厂就是将不同工厂生产的产品关联起来

// 鼠标接口
type IMouse interface {
	SayMouseBrand()
}

// 戴尔鼠标
type DellMouse struct{}

func (d DellMouse) SayMouseBrand() {
	fmt.Println("Dell Mouse")
}

// 惠普鼠标
type HpMouse struct{}

func (h HpMouse) SayMouseBrand() {
	fmt.Println("Hp Mouse")
}

// 键盘接口
type IKeyBoard interface {
	SayKeyBoardBrand()
}

// 戴尔键盘
type DellKeyBoard struct{}

func (d DellKeyBoard) SayKeyBoardBrand() {
	fmt.Println("Dell KeyBoard")
}

// 惠普键盘
type HpKeyBoard struct{}

func (h HpKeyBoard) SayKeyBoardBrand() {
	fmt.Println("Hp KeyBoard")
}

// 抽象工厂
type IPcFactory interface {
	CreateMouse()    //创建鼠标
	CreateKeyBoard() //创建键盘
}

// 惠普工厂，惠普鼠标 + 戴尔键盘
type HpFactory struct {
	Mouse    IMouse
	KeyBoard IKeyBoard
}

func (h *HpFactory) CreateMouse() {
	h.Mouse = HpMouse{}
}

func (h *HpFactory) CreateKeyBoard() {
	h.KeyBoard = DellKeyBoard{}
}

// 戴尔工厂 戴尔鼠标 + 戴尔键盘
type DellFactory struct {
	Mouse    IMouse
	KeyBoard IKeyBoard
}

func (d *DellFactory) CreateMouse() {
	d.Mouse = HpMouse{}
}

func (d *DellFactory) CreateKeyBoard() {
	d.KeyBoard = DellKeyBoard{}
}
