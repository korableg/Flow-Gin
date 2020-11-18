package Node

import (
	"encoding/json"
	"github.com/korableg/mini-gin/Mini/Common"
	"github.com/korableg/mini-gin/Mini/Errors"
	"github.com/korableg/mini-gin/Mini/Messages"
)

type Node struct {
	id       string
	name     string
	messages *Messages.MessageQueue
}

func NewNode(name string) (n *Node, err error) {
	if err = checkName(name); err != nil {
		return
	}
	n = &Node{
		name:     name,
		messages: Messages.NewMessageQueue(),
	}
	return
}

func (n *Node) Name() string {
	return n.name
}

func (n *Node) PushMessage(m *Messages.Message) {
	n.messages.Push(m)
}

func (n *Node) FrontMessage() (m *Messages.Message) {
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
	n.messages = Messages.NewMessageQueue()

	return nil

}

func checkName(name string) error {
	if len(name) == 0 {
		return Errors.ERR_NODE_NAME_ISEMPTY
	}
	if !Common.NameMatchedPattern(name) {
		return Errors.ERR_NODE_NAME_NOT_MATCHED_PATTERN
	}
	return nil
}
