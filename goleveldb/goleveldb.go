package goleveldb

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/korableg/mini-gin/Mini/hub"
	"github.com/korableg/mini-gin/Mini/node"
	"github.com/korableg/mini-gin/Mini/repository"
	"github.com/syndtr/goleveldb/leveldb"
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

func (f *GoLevelDB) NewNodeRepository() repository.NodeRepositoryDB {

	dbPath := fmt.Sprintf("%s%c%s%c%s", f.path, filepath.Separator, "db", filepath.Separator, "nodes")

	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		panic(err)
	}
	n := NewNodeRepository(db)
	return n
}

func (f *GoLevelDB) NewHubRepository() repository.HubRepositoryDB {

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

func (hr *HubRepository) Store(hub *hub.Hub) error {

	hubView := make(map[string]interface{}, 0)
	hubView["name"] = hub.Name()

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, hubView)
	if err != nil {
		return err
	}
	err = hr.db.Put([]byte(hub.Name()), buf.Bytes(), nil)

	return err

}

func (hr *HubRepository) All() ([]*hub.Hub, error) {

	hubs := make([]*hub.Hub, 0, 20)

	iterator := hr.db.NewIterator(nil, nil)
	for iterator.Next() {

		hubView := make(map[string]interface{}, 0)
		reader := bytes.NewReader(iterator.Value())
		err := binary.Read(reader, binary.LittleEndian, hubView)
		if err != nil {
			return nil, err
		}
		hub, err := hub.NewHub(hubView["name"].(string))
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
	db *leveldb.DB
}

func NewNodeRepository(db *leveldb.DB) *NodeRepository {
	nr := NodeRepository{db: db}
	return &nr
}

func (nr *NodeRepository) Store(node *node.Node) error {

	if nodeJSON, err := json.Marshal(node); err != nil {
		return err
	} else if err := nr.db.Put([]byte(node.Name()), nodeJSON, nil); err != nil {
		return err
	}
	return nil

}

func (nr *NodeRepository) All() ([]*node.Node, error) {

	nodes := make([]*node.Node, 0, 20)

	iterator := nr.db.NewIterator(nil, nil)

	for iterator.Next() {

		value := iterator.Value()
		node := &node.Node{}
		if err := json.Unmarshal(value, node); err != nil {
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

func (nr *NodeRepository) Close() error {
	return nr.db.Close()
}
