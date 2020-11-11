package Mini

import (
	"github.com/korableg/mini-gin/Mini/Hub"
	"github.com/korableg/mini-gin/Mini/Messages"
	"github.com/korableg/mini-gin/Mini/Node"
)

func (m *Mini) SendMessage(from *Node.Node, h *Hub.Hub, data []byte) *Messages.Message {
	mes := Messages.NewMessage(from.Name(), data)

	f := func(key, value interface{}) bool {
		node := value.(*Node.Node)
		node.PushMessage(mes)
		return true
	}
	h.RangeNodes(f)

	return mes
}

func (m *Mini) GetMessage(n *Node.Node) *Messages.Message {
	return n.FrontMessage()
}

func (m *Mini) RemoveMessage(n *Node.Node) {
	n.RemoveFrontMessage()
}
