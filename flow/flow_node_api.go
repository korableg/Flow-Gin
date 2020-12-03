package flow

import (
	"github.com/korableg/mini-gin/flow/errs"
	"github.com/korableg/mini-gin/flow/hub"
	"github.com/korableg/mini-gin/flow/node"
)

func (m *Flow) NewNode(name string, careful bool) (*node.Node, error) {

	if nodeExists := m.GetNode(name); nodeExists != nil {
		return nil, errs.ErrNodeIsAlreadyExists
	}

	n, err := node.New(name, careful, m.db)
	if err != nil {
		return nil, err
	}

	if err := m.nodes.Store(n); err != nil {
		return nil, err
	}

	return n, nil

}

func (m *Flow) GetNode(name string) (n *node.Node) {
	if value, ok := m.nodes.Load(name); ok {
		n = value
	}
	return
}

func (m *Flow) GetAllNodes() ([]*node.Node, error) {
	nodes := make([]*node.Node, 0, 20)
	f := func(value *node.Node) error { nodes = append(nodes, value); return nil }
	if err := m.nodes.Range(f); err != nil {
		return nil, err
	}
	return nodes, nil
}

func (m *Flow) DeleteNode(name string) error {
	n := m.GetNode(name)
	if n == nil {
		return nil
	}
	err := n.DeleteMessageDB()
	if err != nil {
		return err
	}
	f := func(hub *hub.Hub) error {
		return hub.DeleteNode(n)
	}
	err = m.hubs.Range(f)
	if err != nil {
		return err
	}
	return m.nodes.Delete(name)
}
