package repo

import "strings"

type DB interface {
	NewNodeRepository() NodeDB
	NewHubRepository() HubDB
}

type NodeDB interface {
	Store(node *Node) error
	All() ([]*Node, error)
	Delete(name string) error
	Close() error
}

type HubDB interface {
	Store(node *Hub) error
	All() ([]*Hub, error)
	Delete(name string) error
	Close() error
}

type Hub struct {
	Name [100]rune
}

func (h *Hub) GetName() string {
	return strings.Trim(string(h.Name[:]), "\000")
}

func (h *Hub) SetName(name string) {
	copy(h.Name[:], []rune(name))
}

type Node struct {
	Name [100]rune
}

func (n *Node) GetName() string {
	return strings.Trim(string(n.Name[:]), "\000")
}

func (n *Node) SetName(name string) {
	copy(n.Name[:], []rune(name))
}
