package flow

import (
	"github.com/korableg/mini-gin/flow/errs"
	"github.com/korableg/mini-gin/flow/msgs"
)

func (m *Flow) SendMessage(from, to string, data []byte) (*msgs.Message, error) {

	h := m.GetHub(to)
	if h == nil {
		return nil, errs.ErrHubNotFound
	}
	n := m.GetNode(from)
	if n == nil {
		return nil, errs.ErrNodeNotFound
	}
	mes := msgs.NewMessage(n.Name(), data)
	h.PushMessage(mes)
	return mes, nil

}

func (m *Flow) GetMessage(nodeName string) (*msgs.Message, error) {
	n := m.GetNode(nodeName)
	if n == nil {
		return nil, errs.ErrNodeNotFound
	}
	return n.FrontMessage(), nil
}

func (m *Flow) RemoveMessage(nodeName string) error {
	n := m.GetNode(nodeName)
	if n == nil {
		return errs.ErrNodeNotFound
	}
	return n.RemoveFrontMessage()
}
