package goleveldb

import (
	"github.com/korableg/mini-gin/Mini/Hub"
	"github.com/syndtr/goleveldb/leveldb"
)

type HubRepository struct {
	db *leveldb.DB
}

func (*HubRepository) Store(key string, value *Hub.Hub) {

}

func (*HubRepository) Load(key string) (*Hub.Hub, bool) {
	return &Hub.Hub{}, true
}

func (*HubRepository) Delete(key string) {

}

func (*HubRepository) Range(f func(key string, value *Hub.Hub) bool) {

}
