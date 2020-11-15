package goleveldb

import (
	"fmt"
	"github.com/korableg/mini-gin/Mini/Hub"
	"github.com/korableg/mini-gin/Mini/Node"
	"github.com/syndtr/goleveldb/leveldb"
	"path/filepath"
	"strings"
)

type Factory struct {
	path string
}

func NewFactory(path string) *Factory {
	if !strings.HasSuffix(path, string(filepath.Separator)) {
		path += string(filepath.Separator)
	}
	f := Factory{
		path: path,
	}
	return &f
}

func (f *Factory) NewNodeRepository() Node.NodeRepositoryDB {

	dbPath := fmt.Sprintf("%s%c%s%c%s", f.path, filepath.Separator, "db", filepath.Separator, "nodes")

	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		panic(err)
	}
	n := NewNodeRepository(db)
	return n
}

func (f *Factory) NewHubRepository() Hub.HubRepository {

	dbPath := fmt.Sprintf("%s%c%s%c%s", f.path, filepath.Separator, "db", filepath.Separator, "hubs")

	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		panic(err)
	}
	n := NewHubRepository(db)
	return n
}
