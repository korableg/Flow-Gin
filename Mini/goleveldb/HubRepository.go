package goleveldb

import (
	"encoding/json"
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

func (hr *HubRepository) Store(key string, value *Hub.Hub) error {

	if hubJSON, err := json.Marshal(value); err != nil {
		return err
	} else if err := hr.db.Put([]byte(key), hubJSON, nil); err != nil {
		return err
	}
	return nil

}

func (hr *HubRepository) All() ([]*Hub.Hub, error) {

	hubs := make([]*Hub.Hub, 0, 20)

	iterator := hr.db.NewIterator(nil, nil)

	for iterator.Next() {

		value := iterator.Value()
		hub := &Hub.Hub{}
		if err := json.Unmarshal(value, hub); err != nil {
			return nil, err
		}

		hubs = append(hubs, hub)

	}

	iterator.Release()

	if err := iterator.Error(); err != nil {
		return nil, err
	}

	return hubs, nil

}

func (hr *HubRepository) Delete(key string) error {
	return hr.db.Delete([]byte(key), nil)
}

func (hr *HubRepository) Close() error {
	return hr.Close()
}
