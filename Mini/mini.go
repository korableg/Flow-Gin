package Mini

import (
	"github.com/korableg/mini-gin/Mini/repository"
)

type Mini struct {
	nodes *repository.NodeRepository
	hubs  *repository.HubRepository
}

func New(factory repository.DB) *Mini {

	var nodeDB repository.NodeRepositoryDB
	var hubDB repository.HubRepositoryDB

	if factory != nil {
		nodeDB = factory.NewNodeRepository()
		hubDB = factory.NewHubRepository()
	}

	m := &Mini{
		nodes: repository.NewNodeRepository(nodeDB),
		hubs:  repository.NewHubRepository(hubDB),
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
