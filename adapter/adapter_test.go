package adapter

import (
	"adapter/plug"
	"testing"
)

func TestExampleToCharge(t *testing.T) {
	c1 := plug.TwoPlug{}
	c2 := plug.ThreePlugAdapter{}
	ToCharge(c1)
	ToCharge(c2)
	//OutPut:
	//我是2孔插头，我使用方法：Charge充电
	//我是3孔插头，我使用方法：ThreePlugCharge充电
}
