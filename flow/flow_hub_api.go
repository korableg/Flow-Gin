package flow

import (
	"github.com/korableg/mini-gin/flow/errs"
	"github.com/korableg/mini-gin/flow/hub"
)

func (m *Flow) NewHub(name string) (h *hub.Hub, err error) {

	if m.GetHub(name) != nil {
		err = errs.ErrHubIsAlreadyExists
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
	hubs := make([]*hub.Hub, 0, 20)

	f := func(hub *hub.Hub) error {
		hubs = append(hubs, hub)
		return nil
	}
	m.hubs.Range(f)

	return hubs
}

func (m *Flow) AddNodeToHub(hubName, nodeName string) error {
	h := m.GetHub(hubName)
	if h == nil {
		return errs.ErrHubNotFound
	}
	n := m.GetNode(nodeName)
	if n == nil {
		return errs.ErrNodeNotFound
	}
	h.AddNode(n)
	return nil
}

func (m *Flow) DeleteNodeFromHub(hubName, nodeName string) error {
	h := m.GetHub(hubName)
	if h == nil {
		return errs.ErrHubNotFound
	}
	n := m.GetNode(nodeName)
	if n == nil {
		return errs.ErrNodeNotFound
	}
	h.DeleteNode(n)
	return nil
}

func (m *Flow) DeleteHub(name string) error {
	return m.hubs.Delete(name)
}
