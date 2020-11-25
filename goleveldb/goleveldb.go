package goleveldb

import (
	"encoding/json"
	"fmt"
	"github.com/korableg/mini-gin/flow/repo"
	"github.com/syndtr/goleveldb/leveldb"
	"os"
	"path/filepath"
	"strings"
)

type GoLevelDB struct {
	path string
}

func New(path string) *GoLevelDB {
	if !strings.HasSuffix(path, string(filepath.Separator)) {
		path += string(filepath.Separator)
	}
	db := GoLevelDB{
		path: path,
	}
	return &db
}

func (f *GoLevelDB) NewNodeRepository(hubName ...string) repo.NodeDB {

	path := "nodes"
	if hubName != nil && len(hubName) > 0 {
		path = fmt.Sprintf("%s%c%s", "nodesinhubs", filepath.Separator, hubName[0])
	}
	dbPath := fmt.Sprintf("%s%c%s%c%s", f.path, filepath.Separator, "db", filepath.Separator, path)

	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		panic(err)
	}
	n := NewNodeRepository(db, dbPath)
	return n
}

func (f *GoLevelDB) NewHubRepository() repo.HubDB {

	dbPath := fmt.Sprintf("%s%c%s%c%s", f.path, filepath.Separator, "db", filepath.Separator, "hubs")

	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		panic(err)
	}
	n := NewHubRepository(db)
	return n
}

type HubRepository struct {
	db *leveldb.DB
}

func NewHubRepository(db *leveldb.DB) *HubRepository {
	hr := HubRepository{db: db}
	return &hr
}

func (hr *HubRepository) Store(hub *repo.Hub) error {

	dataJSON, err := json.Marshal(hub)
	if err != nil {
		return err
	}
	err = hr.db.Put([]byte(hub.Name), dataJSON, nil)

	return err

}

func (hr *HubRepository) All() ([]*repo.Hub, error) {

	hubs := make([]*repo.Hub, 0, 20)

	iterator := hr.db.NewIterator(nil, nil)
	for iterator.Next() {
		hub := new(repo.Hub)
		err := json.Unmarshal(iterator.Value(), hub)
		if err != nil {
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

type NodeRepository struct {
	db   *leveldb.DB
	path string
}

func NewNodeRepository(db *leveldb.DB, path string) *NodeRepository {
	nr := NodeRepository{db: db, path: path}
	return &nr
}

func (nr *NodeRepository) Store(node *repo.Node) error {

	dataJSON, err := json.Marshal(node)
	if err != nil {
		return err
	}
	err = nr.db.Put([]byte(node.Name), dataJSON, nil)

	return err

}

func (nr *NodeRepository) All() ([]*repo.Node, error) {

	nodes := make([]*repo.Node, 0, 20)

	iterator := nr.db.NewIterator(nil, nil)
	for iterator.Next() {
		node := new(repo.Node)
		err := json.Unmarshal(iterator.Value(), node)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}
	iterator.Release()

	if err := iterator.Error(); err != nil {
		return nil, err
	}

	return nodes, nil

}

func (nr *NodeRepository) Delete(key string) error {
	return nr.db.Delete([]byte(key), nil)
}

func (nr *NodeRepository) DeleteDB() (err error) {
	if err = nr.Close(); err == nil {
		err = os.RemoveAll(nr.path)
	}
	return
}

func (nr *NodeRepository) Close() error {
	return nr.db.Close()
}
