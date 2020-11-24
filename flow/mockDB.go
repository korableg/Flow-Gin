package flow

import (
	"github.com/korableg/mini-gin/flow/repo"
)

type mockDB struct{}

func (t *mockDB) NewNodeRepository(hubName ...string) repo.NodeDB {
	n := new(mockNodeRepository)
	return n
}

func (t *mockDB) NewHubRepository() repo.HubDB {
	n := new(mockHubRepository)
	return n
}

type mockHubRepository struct{}

func (hr *mockHubRepository) Store(hub *repo.Hub) error {
	return nil
}

func (hr *mockHubRepository) All() ([]*repo.Hub, error) {
	hubs := make([]*repo.Hub, 0, 20)
	hubs = append(hubs, &repo.Hub{Name: "mockHub"})
	return hubs, nil
}

func (hr *mockHubRepository) Delete(key string) error {
	return nil
}

func (hr *mockHubRepository) Close() error {
	return nil
}

type mockNodeRepository struct{}

func (nr *mockNodeRepository) Store(node *repo.Node) error {
	return nil
}

func (nr *mockNodeRepository) All() ([]*repo.Node, error) {
	nodes := make([]*repo.Node, 0, 20)
	nodes = append(nodes, &repo.Node{Name: "mockNode"})
	return nodes, nil
}

func (nr *mockNodeRepository) Delete(key string) error {
	return nil
}

func (nr *mockNodeRepository) Close() error {
	return nil
}
