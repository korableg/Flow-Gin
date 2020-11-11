package Mini

import (
	"sync"
)

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
