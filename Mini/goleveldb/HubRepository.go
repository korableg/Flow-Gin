package goleveldb

import (
	"github.com/korableg/mini-gin/Mini/Hub"
	"github.com/syndtr/goleveldb/leveldb"
)

type HubRepository struct {
	db *leveldb.DB
}

func NewHubRepository(db *leveldb.DB) *HubRepository {
	hr := HubRepository{db: db}
	return &hr
}

func (*HubRepository) Store(key string, value *Hub.Hub) error {

	return nil

}

func (hr *HubRepository) All() ([]*Hub.Hub, error) {

	nodes := make([]*Hub.Hub, 0, 20)

	return nodes, nil

}

func (hr *HubRepository) Delete(key string) error {
	return nil
}

func (hr *HubRepository) Close() error {
	return hr.Close()
}
