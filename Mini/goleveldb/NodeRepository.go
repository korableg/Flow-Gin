package goleveldb

import (
	"encoding/json"
	"github.com/korableg/mini-gin/Mini/Node"
	"github.com/syndtr/goleveldb/leveldb"
)

type NodeRepository struct {
	db *leveldb.DB
}

func NewNodeRepository(db *leveldb.DB) *NodeRepository {

	nr := NodeRepository{db: db}
	return &nr

}

func NewHubRepository(db *leveldb.DB) *HubRepository {

	nr := HubRepository{db: db}
	return &nr

}

func (nr *NodeRepository) Store(key string, value *Node.Node) error {

	if nodeJSON, err := json.Marshal(value); err != nil {
		return err
	} else if err := nr.db.Put([]byte(key), nodeJSON, nil); err != nil {
		return err
	}

	return nil

}

func (nr *NodeRepository) All() ([]*Node.Node, error) {

	nodes := make([]*Node.Node, 0, 20)

	iterator := nr.db.NewIterator(nil, nil)

	for iterator.Next() {

		value := iterator.Value()
		node := &Node.Node{}
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
