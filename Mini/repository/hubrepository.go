package repository

import (
	"github.com/korableg/mini-gin/Mini/hub"
	"sync"
)

type HubRepository struct {
	hubs *sync.Map
	db   HubRepositoryDB
}

func NewHubRepository(db HubRepositoryDB) *HubRepository {

	hr := HubRepository{
		hubs: new(sync.Map),
		db:   db,
	}

	if hr.db != nil {
		hubs, err := hr.db.All()
		if err != nil {
			panic(err)
		}
		for _, hub := range hubs {
			hr.hubs.Store(hub.Name(), hub)
		}
	}

	return &hr
}

func (hr *HubRepository) Store(hub *hub.Hub) error {
	if hr.db != nil {
		if err := hr.db.Store(hub); err != nil {
			return err
		}
	}
	hr.hubs.Store(hub.Name(), hub)
	return nil
}

func (hr *HubRepository) Load(name string) (*hub.Hub, bool) {
	if node, ok := hr.hubs.Load(name); ok {
		return node.(*hub.Hub), ok
	}
	return nil, false
}

func (hr *HubRepository) Range(f func(hub *hub.Hub)) {
	rangeFunc := func(key, value interface{}) bool {
		f(value.(*hub.Hub))
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
