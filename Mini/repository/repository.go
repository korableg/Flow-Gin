package repository

import (
	"github.com/korableg/mini-gin/Mini/hub"
	"github.com/korableg/mini-gin/Mini/node"
)

type DB interface {
	NewNodeRepository() NodeRepositoryDB
	NewHubRepository() HubRepositoryDB
}

type NodeRepositoryDB interface {
	Store(node *node.Node) error
	All() ([]*node.Node, error)
	Delete(name string) error
	Close() error
}

type HubRepositoryDB interface {
	Store(node *hub.Hub) error
	All() ([]*hub.Hub, error)
	Delete(name string) error
	Close() error
}
