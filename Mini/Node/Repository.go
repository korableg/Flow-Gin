package Node

import (
	"sync"
)

type Repository struct {
	nodes *sync.Map
	db    NodeRepositoryDB
}

func NewRepository(DB NodeRepositoryDB) *Repository {

	nr := Repository{
		nodes: &sync.Map{},
		db:    DB,
	}

	nodes, err := nr.db.All()
	if err != nil {
		panic(err)
	}

	for _, node := range nodes {
		nr.nodes.Store(node.Name(), node)
	}

	return &nr
}

type NodeRepositoryDB interface {
	Store(key string, value *Node) error
	All() ([]*Node, error)
	Delete(key string) error
	Close() error
}

func (nr *Repository) Store(name string, value *Node) error {
	if err := nr.db.Store(name, value); err != nil {
		return err
	}
	nr.nodes.Store(name, value)
	return nil
}

func (nr *Repository) Load(name string) (*Node, bool) {
	if node, ok := nr.nodes.Load(name); ok {
		return node.(*Node), ok
	}
	return nil, false
}

func (nr *Repository) Range(f func(name string, value *Node)) {
	rangeFunc := func(key, value interface{}) bool {
		f(key.(string), value.(*Node))
		return true
	}
	nr.nodes.Range(rangeFunc)
}

func (nr *Repository) Delete(name string) error {
	if err := nr.db.Delete(name); err != nil {
		return err
	}
	nr.nodes.Delete(name)
	return nil
}

func (nr *Repository) Close() error {
	return nr.db.Close()
}
