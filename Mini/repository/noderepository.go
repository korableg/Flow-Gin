package repository

import (
	"github.com/korableg/mini-gin/Mini/node"
	"sync"
)

type NodeRepository struct {
	nodes *sync.Map
	db    NodeRepositoryDB
}

func NewNodeRepository(DB NodeRepositoryDB) *NodeRepository {

	nr := NodeRepository{
		nodes: new(sync.Map),
		db:    DB,
	}

	if nr.db != nil {
		nodes, err := nr.db.All()
		if err != nil {
			panic(err)
		}
		for _, node := range nodes {
			nr.nodes.Store(node.Name(), node)
		}
	}

	return &nr
}

func (nr *NodeRepository) Store(node *node.Node) error {
	if nr.db != nil {
		if err := nr.db.Store(node); err != nil {
			return err
		}
	}
	nr.nodes.Store(node.Name(), node)
	return nil
}

func (nr *NodeRepository) Load(name string) (*node.Node, bool) {
	if n, ok := nr.nodes.Load(name); ok {
		return n.(*node.Node), ok
	}
	return nil, false
}

func (nr *NodeRepository) Range(f func(node *node.Node)) {
	rangeFunc := func(key, value interface{}) bool {
		f(value.(*node.Node))
		return true
	}
	nr.nodes.Range(rangeFunc)
}

func (nr *NodeRepository) Delete(name string) error {
	if nr.db != nil {
		if err := nr.db.Delete(name); err != nil {
			return err
		}
	}
	nr.nodes.Delete(name)
	return nil
}

func (nr *NodeRepository) Close() error {
	if nr.db == nil {
		return nil
	}
	return nr.db.Close()
}
