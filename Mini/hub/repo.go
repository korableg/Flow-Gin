package hub

import (
	"github.com/korableg/mini-gin/Mini/repo"
	"sync"
)

type HubRepository struct {
	hubs *sync.Map
	db   repo.HubDB
}

func NewHubRepository(db repo.HubDB) *HubRepository {

	hr := HubRepository{
		hubs: new(sync.Map),
		db:   db,
	}

	if hr.db != nil {
		hubsRepo, err := hr.db.All()
		if err != nil {
			panic(err)
		}
		for _, hubRepo := range hubsRepo {
			hub, err := New(hubRepo.GetName())
			if err != nil {
				panic(err)
			}
			hr.hubs.Store(hub.Name(), hub)
		}
	}

	return &hr
}

func (hr *HubRepository) Store(hub *Hub) error {
	if hr.db != nil {
		hubRepo := new(repo.Hub)
		hubRepo.SetName(hub.Name())
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
