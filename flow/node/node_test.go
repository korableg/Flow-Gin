package node

import (
	"encoding/json"
	"github.com/korableg/mini-gin/flow/errs"
	"github.com/korableg/mini-gin/flow/msgs"
	"testing"
)

func TestNode(t *testing.T) {

	nameNode := "TestNode1"

	_, err := New("   ")
	if err != errs.ERR_NODE_NAME_NOT_MATCHED_PATTERN {
		t.Error(err)
	}
	_, err = New("")
	if err != errs.ERR_NODE_NAME_ISEMPTY {
		t.Error(err)
	}
	_, err = New("123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")
	if err != errs.ERR_NODE_NAME_OVER100 {
		t.Error(err)
	}
	n, err := New(nameNode)
	if err != nil {
		t.Fatal(err)
	}
	if nameNode != n.Name() {
		t.Error(nameNode + " != " + n.Name())
	}

	n.PushMessage(msgs.NewMessage("TestNode2", nil))

	if n.Len() != 1 {
		t.Error("count messages must be 1")
	}

	m := n.FrontMessage()
	if m == nil {
		t.Error("front message must be not nil")
	}
	n.RemoveFrontMessage()

	if n.Len() != 0 {
		t.Error("count messages must be 0")
	}

	_, err = json.Marshal(n)

}
