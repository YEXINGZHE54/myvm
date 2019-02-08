package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	path string
}

func newDirEntry(path string) *DirEntry {
	e := new(DirEntry)
	abspath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	e.path = abspath
	return e
}

func (e *DirEntry) readClass(clsname string) (buf []byte, ety Entry, err error) {
	path := filepath.Join(e.path, clsname)
	buf, err = ioutil.ReadFile(path)
	ety = e
	return
}

func (e *DirEntry) String() string {
	return e.path
}