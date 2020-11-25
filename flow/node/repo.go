package node

import (
	"github.com/korableg/mini-gin/flow/repo"
	"sync"
)

type NodeRepository struct {
	nodes *sync.Map
	db    repo.NodeDB
}

func NewNodeRepository(db repo.NodeDB) *NodeRepository {

	nr := new(NodeRepository)
	nr.nodes = new(sync.Map)
	nr.db = db

	if nr.db == nil {
		return nr
	}

	nodesRepo, err := nr.db.All()
	if err != nil {
		panic(err)
	}
	for _, nodeRepo := range nodesRepo {
		n, err := New(nodeRepo.Name)
		if err != nil {
			panic(err)
		}
		nr.nodes.Store(n.Name(), n)
	}

	return nr
}

func (nr *NodeRepository) Store(node *Node) error {
	if nr.db != nil {
		nodeRepo := new(repo.Node)
		nodeRepo.Name = node.Name()
		if err := nr.db.Store(nodeRepo); err != nil {
			return err
		}
	}
	nr.nodes.Store(node.Name(), node)
	return nil
}

func (nr *NodeRepository) Load(name string) (*Node, bool) {
	if n, ok := nr.nodes.Load(name); ok {
		return n.(*Node), ok
	}
	return nil, false
}

func (nr *NodeRepository) Range(f func(node *Node)) {
	rangeFunc := func(key, value interface{}) bool {
		f(value.(*Node))
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

func (nr *NodeRepository) DeleteDB() error {
	if nr.db == nil {
		return nil
	}
	err := nr.db.DeleteDB()
	if err != nil {
		return err
	}
	nr.db = nil
	return nil
}

func (nr *NodeRepository) Close() error {
	if nr.db == nil {
		return nil
	}
	return nr.db.Close()
}
