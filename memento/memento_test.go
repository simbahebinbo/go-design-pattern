package memento

import "testing"

func TestExampleGame_Play(t *testing.T) {
	//初始化游戏
	g1 := GameInit()
	//查看游戏状态
	g1.Status()
	//玩游戏
	g1.Play(5, 5)
	//查看游戏状态
	g1.Status()
	//存档
	gm := &GameMemento{}
	gm.Save(g1)
	//初始化游戏
	g2 := GameInit()
	//查看游戏状态
	g2.Status()
	//恢复
	g2 = gm.Load()
	//查看游戏状态
	g2.Status()
	//恢复
	g3 := gm.Load()
	//查看游戏状态
	g3.Status()

	//OutPut:
	//Current HP:0, MP:0
	//Current HP:5, MP:5
	//Current HP:0, MP:0
	//Current HP:5, MP:5
	//Current HP:5, MP:5
}
