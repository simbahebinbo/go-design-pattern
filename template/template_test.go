package template

import "testing"

func TestExampleMan_Dress(t *testing.T) {
	m := Man{p: &Person{}}
	m.SetSex()    //自定义的方法
	m.p.Eat("米饭") //使用模板中的方法
	m.Dress()     //自定义的方法
	//OutPut:
	//吃食物  米饭
	//男人 化妆,  刮胡子
}

func TestExampleWomen_Dress(t *testing.T) {
	w := Women{p: &Person{}}
	w.SetSex()     //自定义的方法
	w.p.Eat("麻辣烫") //使用模板中的方法
	w.Dress()      //自定义的方法
	//OutPut:
	//吃食物  麻辣烫
	//女人 化妆,  涂口红
}
