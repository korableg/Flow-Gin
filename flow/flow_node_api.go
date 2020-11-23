package flow

import (
	"github.com/korableg/mini-gin/flow/errs"
	"github.com/korableg/mini-gin/flow/hub"
	"github.com/korableg/mini-gin/flow/node"
)

func (m *Flow) NewNode(name string) (n *node.Node, err error) {

	if nodeExists := m.GetNode(name); nodeExists != nil {
		err = errs.ERR_NODE_IS_ALREADY_EXISTS
		return
	}

	n, err = node.NewNode(name)
	if err != nil {
		return
	}

	m.nodes.Store(n)
	return

}

func (m *Flow) GetNode(name string) (n *node.Node) {
	if value, ok := m.nodes.Load(name); ok {
		n = value
	}
	return
}

func (m *Flow) GetAllNodes() []*node.Node {
	nodes := make([]*node.Node, 0, 20)

	f := func(value *node.Node) {
		nodes = append(nodes, value)
	}
	m.nodes.Range(f)

	return nodes
}

func (m *Flow) DeleteNode(name string) {

	node := m.GetNode(name)
	if node == nil {
		return
	}

	f := func(hub *hub.Hub) {
		hub.DeleteNode(node)
	}

	m.hubs.Range(f)

	m.nodes.Delete(name)

}
