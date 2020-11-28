package hub

import (
	"encoding/json"
	"github.com/korableg/mini-gin/flow/cmn"
	"github.com/korableg/mini-gin/flow/errs"
	"github.com/korableg/mini-gin/flow/msgs"
	"github.com/korableg/mini-gin/flow/node"
	"github.com/korableg/mini-gin/flow/repo"
)

type Hub struct {
	name  string
	nodes *node.NodeRepository
}

func New(name string, db repo.DB) (h *Hub, err error) {

	if err = checkName(name); err != nil {
		return
	}

	var nodeDB repo.NodeDB
	if db != nil {
		nodeDB = db.NewNodeRepository(name)
	}

	h = &Hub{
		name:  name,
		nodes: node.NewNodeRepository(nodeDB),
	}

	return

}

func (h *Hub) Name() string {
	return h.name
}

func (h *Hub) AddNode(n *node.Node) error {
	return h.nodes.Store(n)
}

func (h *Hub) DeleteNode(n *node.Node) error {
	if n == nil {
		return nil
	}
	return h.nodes.Delete(n.Name())
}

func (h *Hub) PushMessage(m *msgs.Message) {
	h.rangeNodes(func(n *node.Node) { n.PushMessage(m) })
}

func (h *Hub) MarshalJSON() ([]byte, error) {

	nodes := make([]*node.Node, 0, 20)

	f := func(n *node.Node) {
		nodes = append(nodes, n)
	}
	h.rangeNodes(f)

	hubMap := make(map[string]interface{})
	hubMap["name"] = h.name
	hubMap["nodes"] = nodes

	return json.Marshal(hubMap)

}

func (h *Hub) rangeNodes(f func(n *node.Node)) {
	h.nodes.Range(f)
}

func (h *Hub) deleteAllNodes() (err error) {
	nodes := make([]*node.Node, 0, 20)
	h.rangeNodes(func(n *node.Node) { nodes = append(nodes, n) })
	for _, n := range nodes {
		err = h.DeleteNode(n)
		if err != nil {
			return
		}
	}
	return
}

func (h *Hub) deleteNodeDB() error {
	err := h.nodes.DeleteDB()
	return err
}

func checkName(name string) error {
	if len(name) == 0 {
		return errs.ERR_HUB_NAME_ISEMPTY
	}
	if len([]rune(name)) > 100 {
		return errs.ERR_HUB_NAME_OVER100
	}
	if !cmn.NameMatchedPattern(name) {
		return errs.ERR_HUB_NAME_NOT_MATCHED_PATTERN
	}
	return nil
}
