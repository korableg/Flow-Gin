package repo

type DB interface {
	NewNodeRepository(hubName ...string) NodeDB
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

type Node struct {
	Name string
}

type Hub struct {
	Name  string
	Nodes []Node
}
