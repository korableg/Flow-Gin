package Mini

import (
	"github.com/korableg/mini-gin/Mini/Errors"
	"github.com/korableg/mini-gin/Mini/Hub"
	"github.com/korableg/mini-gin/Mini/Node"
)

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

func (m *Mini) GetAllHubs() []*Hub.Hub {
	hubs := make([]*Hub.Hub, 0)

	f := func(key, value interface{}) bool {
		hubs = append(hubs, value.(*Hub.Hub))
		return true
	}
	m.hubs.Range(f)

	return hubs
}

func (m *Mini) AddNodeToHub(h *Hub.Hub, n *Node.Node) {
	h.AddNode(n)
}

func (m *Mini) DeleteNodeFromHub(h *Hub.Hub, n *Node.Node) {
	h.DeleteNode(n)
}

func (m *Mini) DeleteHub(name string) {
	m.hubs.Delete(name)
}
