package Mini

import (
	"github.com/korableg/mini-gin/Mini/Hub"
	"github.com/korableg/mini-gin/Mini/Node"
)

type Mini struct {
	nodes *Node.NodeRepository
	hubs  *Hub.HubRepository
}

type RepositoryDBFactory interface {
	NewNodeRepository() Node.NodeRepositoryDB
	NewHubRepository() Hub.HubRepositoryDB
}

func NewMini(factory RepositoryDBFactory) *Mini {

	m := &Mini{
		nodes: Node.NewNodeRepository(factory.NewNodeRepository),
		hubs:  Hub.NewHubRepository(factory.NewHubRepository),
	}

	return m

}

func (m *Mini) Close() error {
	if err := m.nodes.Close(); err != nil {
		return err
	}

	if err := m.hubs.Close(); err != nil {
		return err
	}

	return nil
}
