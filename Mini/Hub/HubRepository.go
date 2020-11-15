package Hub

type HubRepository interface {
	Store(key string, value *Hub)
	Load(key string) (*Hub, bool)
	Delete(key string)
	Range(f func(key string, value *Hub) bool)
}
