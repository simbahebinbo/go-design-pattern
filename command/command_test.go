package command

import "testing"

func TestExampleInvoker_Do(t *testing.T) {
	//电视实例
	tv1 := &TV{Name: "熊猫"}
	//实例化打开命令
	opencmd1 := OpenCommand{Tv: tv1}
	//执行打开命令
	invoker1 := new(Invoker)
	invoker1.SetCommand(opencmd1)
	invoker1.Do()
	//实例化关闭命令
	closecmd1 := CloseCommand{Tv: tv1}
	//执行关闭命令
	invoker1.SetCommand(closecmd1)
	invoker1.Do()

	//更换电视实例
	tv2 := &TV{Name: "黑熊"}
	opencmd2 := OpenCommand{Tv: tv2}
	invoker2 := &Invoker{}
	invoker2.SetCommand(opencmd2)
	invoker2.Do()

	//OutPut:
	//打开熊猫电视
	//关闭熊猫电视
	//打开黑熊电视
}
