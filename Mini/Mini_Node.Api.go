package Mini

import (
	"github.com/korableg/mini-gin/Mini/Errors"
	"github.com/korableg/mini-gin/Mini/Hub"
	"github.com/korableg/mini-gin/Mini/Node"
)

func (m *Mini) NewNode(name string) (n *Node.Node, err error) {

	if nodeExists := m.GetNode(name); nodeExists != nil {
		err = Errors.ERR_NODE_IS_ALREADY_EXISTS
		return
	}

	n, err = Node.NewNode(name)
	if err != nil {
		return
	}

	m.nodes.Store(n.Name(), n)
	return

}

func (m *Mini) GetNode(name string) (n *Node.Node) {
	if value, ok := m.nodes.Load(name); ok {
		n = value.(*Node.Node)
	}
	return
}

func (m *Mini) GetAllNodes() []*Node.Node {
	nodes := make([]*Node.Node, 0)

	f := func(key, value interface{}) bool {
		nodes = append(nodes, value.(*Node.Node))
		return true
	}
	m.nodes.Range(f)

	return nodes
}

func (m *Mini) DeleteNode(name string) {

	node := m.GetNode(name)
	if node == nil {
		return
	}

	m.nodes.Delete(name)

	f := func(key, value interface{}) bool {
		hub := value.(*Hub.Hub)
		hub.DeleteNode(node)
		return true
	}

	m.hubs.Range(f)

}
