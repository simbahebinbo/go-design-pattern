package bridge

import "testing"

func TestExampleCommonMessageViaSMS(t *testing.T) {
	m := NewCommonMessage(ViaSMS())
	m.SendMessage("have a drink?", "bob")
	//OutPut:
	//send [Common] have a drink? to bob via SMS
}

func TestExampleUrgencyMessageViaSMS(t *testing.T) {
	m := NewUrgencyMessage(ViaSMS())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send [Urgency] have a drink? to bob via SMS
}

func TestExampleCommonMessageViaEmail(t *testing.T) {
	m := NewCommonMessage(ViaEmail())
	m.SendMessage("have a drink?", "bob")
	//OutPut:
	//send [Common] have a drink? to bob via Email
}

func TestExampleUrgencyMessageViaEmail(t *testing.T) {
	m := NewUrgencyMessage(ViaEmail())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send [Urgency] have a drink? to bob via Email
}
