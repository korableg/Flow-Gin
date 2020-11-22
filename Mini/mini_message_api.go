package Mini

import (
	"github.com/korableg/mini-gin/Mini/hub"
	"github.com/korableg/mini-gin/Mini/msgs"
	"github.com/korableg/mini-gin/Mini/node"
)

func (m *Mini) SendMessage(from *node.Node, h *hub.Hub, data []byte) *msgs.Message {
	mes := msgs.NewMessage(from.Name(), data)

	f := func(key, value interface{}) bool {
		node := value.(*node.Node)
		node.PushMessage(mes)
		return true
	}
	h.RangeNodes(f)

	return mes
}

func (m *Mini) GetMessage(n *node.Node) *msgs.Message {
	return n.FrontMessage()
}

func (m *Mini) RemoveMessage(n *node.Node) {
	n.RemoveFrontMessage()
}
