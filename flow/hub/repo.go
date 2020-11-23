package hub

import (
	"github.com/korableg/mini-gin/flow/node"
	"github.com/korableg/mini-gin/flow/repo"
	"sync"
)

type HubRepository struct {
	hubs *sync.Map
	db   repo.HubDB
}

func NewHubRepository(db repo.DB, nodes *node.NodeRepository) *HubRepository {

	hr := new(HubRepository)
	hr.hubs = new(sync.Map)

	if db == nil {
		return hr
	}

	hr.db = db.NewHubRepository()
	hubsRepo, err := hr.db.All()
	if err != nil {
		panic(err)
	}
	for _, hubRepo := range hubsRepo {
		hub, err := New(hubRepo.Name, db)
		if err != nil {
			panic(err)
		}
		for _, nodeRepo := range hubRepo.Nodes {
			if n, ok := nodes.Load(nodeRepo.Name); ok {
				hub.AddNode(n)
			}
		}
		hr.hubs.Store(hub.Name(), hub)
	}

	return hr
}

func (hr *HubRepository) Store(hub *Hub) error {
	if hr.db != nil {
		hubRepo := new(repo.Hub)
		hubRepo.Name = hub.Name()
		if err := hr.db.Store(hubRepo); err != nil {
			return err
		}
	}
	hr.hubs.Store(hub.Name(), hub)
	return nil
}

func (hr *HubRepository) Load(name string) (*Hub, bool) {
	if node, ok := hr.hubs.Load(name); ok {
		return node.(*Hub), ok
	}
	return nil, false
}

func (hr *HubRepository) Range(f func(hub *Hub)) {
	rangeFunc := func(key, value interface{}) bool {
		f(value.(*Hub))
		return true
	}
	hr.hubs.Range(rangeFunc)
}

func (hr *HubRepository) Delete(name string) error {
	if hr.db != nil {
		if err := hr.db.Delete(name); err != nil {
			return err
		}
	}
	hr.hubs.Delete(name)
	return nil
}

func (hr *HubRepository) Close() error {
	if hr.db == nil {
		return nil
	}
	return hr.db.Close()
}
