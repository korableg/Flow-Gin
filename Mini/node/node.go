package node

import (
	"encoding/json"
	"github.com/korableg/mini-gin/Mini/cmn"
	"github.com/korableg/mini-gin/Mini/errs"
	"github.com/korableg/mini-gin/Mini/msgs"
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

func (n *Node) UnmarshalJSON(data []byte) error {

	nodeMap := make(map[string]interface{})
	if err := json.Unmarshal(data, &nodeMap); err != nil {
		return err
	}

	var nodeName string

	if name := nodeMap["name"]; name != nil {
		nodeName = name.(string)
	}

	if err := checkName(nodeName); err != nil {
		return err
	}

	n.name = nodeName
	n.messages = msgs.NewMessageQueue()

	return nil

}

func checkName(name string) error {
	if len(name) == 0 {
		return errs.ERR_NODE_NAME_ISEMPTY
	}
	if !cmn.NameMatchedPattern(name) {
		return errs.ERR_NODE_NAME_NOT_MATCHED_PATTERN
	}
	return nil
}
