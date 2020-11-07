package Mini

import "sync"

var instance *Mini

func init() {
	instance = NewMini()
}

func GetMini() *Mini {
	return instance
}

type Mini struct {
	nodes *sync.Map
	hubs  *sync.Map
}

func NewMini() *Mini {

	m := &Mini{
		nodes: &sync.Map{},
		hubs:  &sync.Map{},
	}
	return m

}
