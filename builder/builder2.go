// 建造者模式
package builder

//意图：将一个复杂的构建与其表示相分离，使得同样的构建过程可以创建不同的表示
//解决:复杂对象的各个部分经常面临着剧烈的变化，但是将它们组合在一起的算法却相对稳定

//实例:组合一个大门，有锁，把手，门，不同的材料不同的价格

//组件不变，组合经常变化

type Item interface {
	Name() string
	Price() float32
}

// 锁
type Lock struct {
}

func (Lock) Name() string {
	return "锁"
}
func (Lock) Price() float32 {
	return 22.22
}

// 把手
type Handle struct {
}

func (Handle) Name() string {
	return "把手"
}
func (Handle) Price() float32 {
	return 11.11
}

// 门
type Door struct {
}

func (Door) Name() string {
	return "门"
}
func (Door) Price() float32 {
	return 33.33
}

type BigDoor []Item

func (bd *BigDoor) AddItem(item ...Item) {
	*bd = append(*bd, item...)
}

func (bd BigDoor) GetCost() (cost float32) {
	for _, v := range bd {
		cost += v.Price()
	}
	return
}

func (bd BigDoor) ShowItems() (msg string) {
	for _, v := range bd {
		msg += "  组件:" + v.Name()
	}
	return
}

// 建造
type DoorBuilder struct {
}

func (DoorBuilder) Build(item ...Item) *BigDoor {
	bigdoor := new(BigDoor)
	bigdoor.AddItem(item...)
	return bigdoor
}
