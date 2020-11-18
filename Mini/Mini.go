package Mini

import (
	"github.com/korableg/mini-gin/Mini/Hub"
	"github.com/korableg/mini-gin/Mini/Node"
)

type Mini struct {
	nodes *Node.Repository
	hubs  *Hub.Repository
}

type RepositoryDBFactory interface {
	NewNodeRepository() Node.NodeRepositoryDB
	NewHubRepository() Hub.HubRepositoryDB
}

func NewMini(factory RepositoryDBFactory) *Mini {

	nodeDB := factory.NewNodeRepository()
	hubDB := factory.NewHubRepository()

	m := &Mini{
		nodes: Node.NewRepository(nodeDB),
	}

	m.hubs = Hub.NewRepository(hubDB, nodeDB, m.nodes)

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
