package node

import (
	"encoding/json"
	"github.com/korableg/mini-gin/flow/cmn"
	"github.com/korableg/mini-gin/flow/errs"
	"github.com/korableg/mini-gin/flow/msgs"
)

type Node struct {
	id       string
	name     string
	messages *msgs.MessageQueue
}

func NewNode(name string) (n *Node, err error) {
	if err = checkName(name); err != nil {
		return
	}
	n = &Node{
		name:     name,
		messages: msgs.NewMessageQueue(),
	}
	return
}

func (n *Node) Name() string {
	return n.name
}

func (n *Node) PushMessage(m *msgs.Message) {
	n.messages.Push(m)
}

func (n *Node) FrontMessage() (m *msgs.Message) {
	m = n.messages.Front()
	return
}

func (n *Node) RemoveFrontMessage() {
	n.messages.RemoveFront()
}

func (n *Node) Len() int {
	return n.messages.Len()
}

func (n *Node) MarshalJSON() ([]byte, error) {

	nodeMap := make(map[string]interface{})
	nodeMap["name"] = n.name

	return json.Marshal(nodeMap)

}

func checkName(name string) error {
	if len(name) == 0 {
		return errs.ERR_NODE_NAME_ISEMPTY
	}
	if len([]rune(name)) > 100 {
		return errs.ERR_NODE_NAME_OVER100
	}
	if !cmn.NameMatchedPattern(name) {
		return errs.ERR_NODE_NAME_NOT_MATCHED_PATTERN
	}
	return nil
}
