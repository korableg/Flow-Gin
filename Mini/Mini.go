package Mini

import (
	"github.com/korableg/mini-gin/Mini/Errors"
	"github.com/korableg/mini-gin/Mini/Hub"
	"github.com/korableg/mini-gin/Mini/Messages"
	"github.com/korableg/mini-gin/Mini/Node"
	"sync"
)

type Mini struct {
	nodes *sync.Map
	hubs  *sync.Map
}

func NewMini() *Mini {

	m := &Mini{
		nodes: &sync.Map{},
		hubs:  &sync.Map{},
	}
	return m

}

func (m *Mini) NewHub(name string) (h *Hub.Hub, err error) {

	if m.GetHub(name) != nil {
		err = Errors.ERR_HUB_IS_ALREADY_EXISTS
		return
	}

	h, err = Hub.NewHub(name)
	m.hubs.Store(name, h)

	return
}

func (m *Mini) GetHub(name string) (h *Hub.Hub) {
	if value, ok := m.hubs.Load(name); ok {
		h = value.(*Hub.Hub)
	}
	return
}

func (m *Mini) AddNodeToHub(h *Hub.Hub, n *Node.Node) {
	h.AddNode(n)
}

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
