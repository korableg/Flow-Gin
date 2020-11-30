package node

import (
	"encoding/json"
	"github.com/korableg/mini-gin/flow/errs"
	"github.com/korableg/mini-gin/flow/msgs"
	MockDB "github.com/korableg/mini-gin/flow/repo/mockDB"
	"testing"
)

func TestNode(t *testing.T) {

	db := new(MockDB.MockDB)

	nameNode := "TestNode1"
	_, err := New("   ", true, nil)
	if err != errs.ERR_NODE_NAME_NOT_MATCHED_PATTERN {
		t.Error(err)
	}
	_, err = New("", false, db)
	if err != errs.ERR_NODE_NAME_ISEMPTY {
		t.Error(err)
	}
	_, err = New("123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890", true, db)
	if err != errs.ERR_NODE_NAME_OVER100 {
		t.Error(err)
	}
	n, err := New(nameNode, false, db)
	if err != nil {
		t.Fatal(err)
	}
	if n.IsCareful() {
		t.Error("node must be not careful")
	}
	if nameNode != n.Name() {
		t.Error(nameNode + " != " + n.Name())
	}

	for n.Len() > 0 {
		n.RemoveFrontMessage()
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

	err = n.DeleteMessageDB()
	if err != nil {
		t.Error(err)
	}

	_, err = json.Marshal(n)

}
