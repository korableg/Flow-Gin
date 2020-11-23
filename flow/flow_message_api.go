package flow

import (
	"github.com/korableg/mini-gin/flow/hub"
	"github.com/korableg/mini-gin/flow/msgs"
	"github.com/korableg/mini-gin/flow/node"
)

func (m *Flow) SendMessage(from *node.Node, h *hub.Hub, data []byte) *msgs.Message {

	mes := msgs.NewMessage(from.Name(), data)
	h.RangeNodes(func(n *node.Node) { n.PushMessage(mes) })

	return mes
}

func (m *Flow) GetMessage(n *node.Node) *msgs.Message {
	return n.FrontMessage()
}

func (m *Flow) RemoveMessage(n *node.Node) {
	n.RemoveFrontMessage()
}
