package hub

import (
	"encoding/json"
	"github.com/korableg/mini-gin/flow/cmn"
	"github.com/korableg/mini-gin/flow/errs"
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

func (h *Hub) AddNode(n *node.Node) {
	h.nodes.Store(n)
}

func (h *Hub) DeleteNode(n *node.Node) {
	h.nodes.Delete(n.Name())
}

func (h *Hub) RangeNodes(f func(n *node.Node)) {
	h.nodes.Range(f)
}

func (h *Hub) MarshalJSON() ([]byte, error) {

	nodes := make([]*node.Node, 0, 20)

	f := func(n *node.Node) {
		nodes = append(nodes, n)
	}
	h.RangeNodes(f)

	hubMap := make(map[string]interface{})
	hubMap["name"] = h.name
	hubMap["nodes"] = nodes

	return json.Marshal(hubMap)

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
