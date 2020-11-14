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

	if len(name) == 0 {
		err = Errors.ERR_NODE_NAME_ISEMPTY
		return
	}

	if !Common.NameMatchedPattern(name) {
		err = Errors.ERR_NODE_NAME_NOT_MATCHED_PATTERN
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
	nodeMap["messages"] = n.Len()

	return json.Marshal(nodeMap)

}

//func (n *Node) UnmarshalJSON(data []byte) error {
//
//	nodeMap := make(map[string]interface{})
//	if err := json.Unmarshal(data, &nodeMap); err != nil {
//		return err
//	}
//	if name := nodeMap["name"]; name != nil {
//		n.name = name.(string)
//	} else {
//		return Errors.ERR_NODE_NAME_ISEMPTY
//	}
//	if n.messages == nil {
//		n.messages = Messages.NewMessageQueue()
//	}
//	return nil
//
//}
