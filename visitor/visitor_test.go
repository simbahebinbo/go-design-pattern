package visitor

import "testing"

func TestExampleObjectStructure_Accept(t *testing.T) {
	object := ObjectStructure{}
	object.Attach(&ConcreteElementA{"A"})
	object.Attach(&ConcreteElementB{"B"})

	cva := ConcreteVisitorA{"vA"}
	cvb := ConcreteVisitorB{"vB"}

	object.Accept(&cva)
	object.Accept(&cvb)

	//OutPut:
	//A vA
	//OperatorA
	//B vA
	//OperatorB
	//A vB
	//OperatorA
	//B vB
	//OperatorB
}
