package hub

import (
	"encoding/json"
	"github.com/korableg/mini-gin/Mini/cmn"
	"github.com/korableg/mini-gin/Mini/errs"
	"github.com/korableg/mini-gin/Mini/node"
	"sync"
)

type Hub struct {
	name  string
	nodes *sync.Map
}

func New(name string) (h *Hub, err error) {

	if err = checkName(name); err != nil {
		return
	}

	h = &Hub{
		name:  name,
		nodes: &sync.Map{},
	}

	return

}

func (h *Hub) Name() string {
	return h.name
}

func (h *Hub) AddNode(n *node.Node) {
	h.nodes.Store(n.Name(), n)
}

func (h *Hub) DeleteNode(n *node.Node) {
	h.nodes.Delete(n.Name())
}

func (h *Hub) RangeNodes(f func(key, value interface{}) bool) {
	h.nodes.Range(f)
}

func (h *Hub) MarshalJSON() ([]byte, error) {

	nodes := make([]*node.Node, 0, 20)

	f := func(key, value interface{}) bool {
		node := value.(*node.Node)
		nodes = append(nodes, node)
		return true
	}
	h.RangeNodes(f)

	hubMap := make(map[string]interface{})
	hubMap["name"] = h.name
	hubMap["nodes"] = nodes

	return json.Marshal(hubMap)

}

func (h *Hub) UnmarshalJSON(data []byte) error {

	hubMap := make(map[string]interface{})
	if err := json.Unmarshal(data, &hubMap); err != nil {
		return err
	}

	var hubName string

	if name := hubMap["name"]; name != nil {
		hubName = name.(string)
	}

	if err := checkName(hubName); err != nil {
		return err
	}

	h.name = hubName
	h.nodes = &sync.Map{}

	return nil

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
