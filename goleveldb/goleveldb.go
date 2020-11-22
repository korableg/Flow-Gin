package goleveldb

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/korableg/mini-gin/Mini/repo"
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

func (f *GoLevelDB) NewNodeRepository() repo.NodeDB {

	dbPath := fmt.Sprintf("%s%c%s%c%s", f.path, filepath.Separator, "db", filepath.Separator, "nodes")

	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		panic(err)
	}
	n := NewNodeRepository(db)
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

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, hub)
	if err != nil {
		return err
	}
	err = hr.db.Put([]byte(hub.GetName()), buf.Bytes(), nil)

	return err

}

func (hr *HubRepository) All() ([]*repo.Hub, error) {

	hubs := make([]*repo.Hub, 0, 20)

	iterator := hr.db.NewIterator(nil, nil)
	for iterator.Next() {
		hub := new(repo.Hub)
		reader := bytes.NewReader(iterator.Value())
		err := binary.Read(reader, binary.LittleEndian, hub)
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
	db *leveldb.DB
}

func NewNodeRepository(db *leveldb.DB) *NodeRepository {
	nr := NodeRepository{db: db}
	return &nr
}

func (nr *NodeRepository) Store(node *repo.Node) error {

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, node)
	if err != nil {
		return err
	}
	err = nr.db.Put([]byte(node.GetName()), buf.Bytes(), nil)

	return err

}

func (nr *NodeRepository) All() ([]*repo.Node, error) {

	nodes := make([]*repo.Node, 0, 20)

	iterator := nr.db.NewIterator(nil, nil)
	for iterator.Next() {
		node := new(repo.Node)
		reader := bytes.NewReader(iterator.Value())
		err := binary.Read(reader, binary.LittleEndian, node)
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

func (nr *NodeRepository) Close() error {
	return nr.db.Close()
}
