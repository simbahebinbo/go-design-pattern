package iterator

import "testing"

func TestExamplePeoplesIterator_Iterator(t *testing.T) {
	p1 := &Person{"小明01", 16, "足球"}
	p2 := &Person{"小明02", 17, "篮球，足球"}
	p3 := &Person{"小明03", 18, "乒乓球"}
	ps := &Peoples{}
	ps.Add(p1, p2, p3)
	pi := ps.CreateIterator()
	pi.Iterator()
	//OutPut:
	//姓名:小明01,年龄:16,爱好:足球
	//姓名:小明02,年龄:17,爱好:篮球，足球
}
