package Hub

import (
	"encoding/json"
	"github.com/korableg/mini-gin/Mini/Common"
	"github.com/korableg/mini-gin/Mini/Errors"
	"github.com/korableg/mini-gin/Mini/Node"
	"sync"
)

type Hub struct {
	name  string
	nodes *sync.Map
}

func NewHub(name string) (h *Hub, err error) {

	if len(name) == 0 {
		err = Errors.ERR_HUB_NAME_ISEMPTY
		return
	}

	if !Common.NameMatchedPattern(name) {
		err = Errors.ERR_HUB_NAME_NOT_MATCHED_PATTERN
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

func (h *Hub) AddNode(n *Node.Node) {
	h.nodes.Store(n.Name(), n)
}

func (h *Hub) DeleteNode(n *Node.Node) {
	h.nodes.Delete(n.Name())
}

func (h *Hub) RangeNodes(f func(key, value interface{}) bool) {
	h.nodes.Range(f)
}

func (h *Hub) MarshalJSON() ([]byte, error) {
	return json.Marshal(h)
}

func (h *Hub) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, h); err != nil {
		return err
	}
	if h.nodes == nil {
		h.nodes = &sync.Map{}
	}
	return nil
}
