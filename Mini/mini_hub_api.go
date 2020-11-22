package Mini

import (
	"github.com/korableg/mini-gin/Mini/errs"
	"github.com/korableg/mini-gin/Mini/hub"
	"github.com/korableg/mini-gin/Mini/node"
)

func (m *Mini) NewHub(name string) (h *hub.Hub, err error) {

	if m.GetHub(name) != nil {
		err = errs.ERR_HUB_IS_ALREADY_EXISTS
		return
	}

	h, err = hub.New(name)
	err = m.hubs.Store(h)

	return
}

func (m *Mini) GetHub(name string) (h *hub.Hub) {
	if value, ok := m.hubs.Load(name); ok {
		h = value
	}
	return
}

func (m *Mini) GetAllHubs() []*hub.Hub {
	hubs := make([]*hub.Hub, 0)

	f := func(hub *hub.Hub) {
		hubs = append(hubs, hub)
	}
	m.hubs.Range(f)

	return hubs
}

func (m *Mini) AddNodeToHub(h *hub.Hub, n *node.Node) {
	h.AddNode(n)
}

func (m *Mini) DeleteNodeFromHub(h *hub.Hub, n *node.Node) {
	h.DeleteNode(n)
}

func (m *Mini) DeleteHub(name string) {
	m.hubs.Delete(name)
}
