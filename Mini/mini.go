package Mini

import (
	"github.com/korableg/mini-gin/Mini/hub"
	"github.com/korableg/mini-gin/Mini/node"
	"github.com/korableg/mini-gin/Mini/repo"
)

type Mini struct {
	nodes *node.NodeRepository
	hubs  *hub.HubRepository
}

func New(factory repo.DB) *Mini {

	var nodeDB repo.NodeDB
	var hubDB repo.HubDB

	if factory != nil {
		nodeDB = factory.NewNodeRepository()
		hubDB = factory.NewHubRepository()
	}

	m := &Mini{
		nodes: node.NewNodeRepository(nodeDB),
		hubs:  hub.NewHubRepository(hubDB),
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
