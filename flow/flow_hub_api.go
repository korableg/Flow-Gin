package flow

import (
	"github.com/korableg/mini-gin/flow/errs"
	"github.com/korableg/mini-gin/flow/hub"
	"github.com/korableg/mini-gin/flow/node"
)

func (m *Flow) NewHub(name string) (h *hub.Hub, err error) {

	if m.GetHub(name) != nil {
		err = errs.ERR_HUB_IS_ALREADY_EXISTS
		return
	}

	h, err = hub.New(name, m.db)
	err = m.hubs.Store(h)

	return
}

func (m *Flow) GetHub(name string) (h *hub.Hub) {
	if value, ok := m.hubs.Load(name); ok {
		h = value
	}
	return
}

func (m *Flow) GetAllHubs() []*hub.Hub {
	hubs := make([]*hub.Hub, 0)

	f := func(hub *hub.Hub) {
		hubs = append(hubs, hub)
	}
	m.hubs.Range(f)

	return hubs
}

func (m *Flow) AddNodeToHub(h *hub.Hub, n *node.Node) {
	h.AddNode(n)
}

func (m *Flow) DeleteNodeFromHub(h *hub.Hub, n *node.Node) {
	h.DeleteNode(n)
}

func (m *Flow) DeleteHub(name string) {
	m.hubs.Delete(name)
}
