package Hub

import "sync"

type HubRepository struct {
	hubs *sync.Map
	db   HubRepositoryDB
}

func NewHubRepository(f func() HubRepositoryDB) *HubRepository {

	nr := HubRepository{
		hubs: &sync.Map{},
		db:   f(),
	}

	hubs, err := nr.db.All()
	if err != nil {
		panic(err)
	}

	for _, hub := range hubs {
		nr.hubs.Store(hub.Name(), hub)
	}

	return &nr
}

type HubRepositoryDB interface {
	Store(key string, value *Hub) error
	All() ([]*Hub, error)
	Delete(key string) error
	Close() error
}

func (nr *HubRepository) Store(name string, value *Hub) error {
	if err := nr.db.Store(name, value); err != nil {
		return err
	}
	nr.hubs.Store(name, value)
	return nil
}

func (nr *HubRepository) Load(name string) (*Hub, bool) {
	if node, ok := nr.hubs.Load(name); ok {
		return node.(*Hub), ok
	}
	return nil, false
}

func (nr *HubRepository) Range(f func(name string, value *Hub)) {
	rangeFunc := func(key, value interface{}) bool {
		f(key.(string), value.(*Hub))
		return true
	}
	nr.hubs.Range(rangeFunc)
}

func (nr *HubRepository) Delete(name string) error {
	if err := nr.db.Delete(name); err != nil {
		return err
	}
	nr.hubs.Delete(name)
	return nil
}

func (nr *HubRepository) Close() error {
	return nr.db.Close()
}
