package Mini

import (
	"github.com/korableg/mini-gin/Mini/Hub"
	"github.com/korableg/mini-gin/Mini/Node"
	"sync"
)

type Mini struct {
	nodes *Node.NodeRepository
	hubs  *sync.Map
}

type RepositoryDBFactory interface {
	NewNodeRepository() Node.NodeRepositoryDB
	NewHubRepository() Hub.HubRepository
}

func NewMini(factory RepositoryDBFactory) *Mini {

	m := &Mini{
		nodes: Node.NewNodeRepository(factory.NewNodeRepository),
		hubs:  &sync.Map{},
	}

	return m

}

func (m *Mini) Close() {
	m.nodes.Close()
}
