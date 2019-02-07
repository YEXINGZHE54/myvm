package classpath

import (
	"strings"
)

type CompositeEntry struct {
	entrys []Entry
}

func newCompositeEntry(path string) Entry {
	e := new(CompositeEntry)
	for _, path := range strings.Split(path, pathListSeparator) {
		e.entrys = append(e.entrys, newEntry(path))
	}
	return e
}

func (e *CompositeEntry) readClass(clsname string) (buf []byte, ety Entry, err error) {
	for _, item := range e.entrys {
		buf, ety, err = item.readClass(clsname)
		if err == nil {
			return
		}
	}
	return nil, nil, ErrorClassNotFound
}

func (e *CompositeEntry) String() string {
	paths := make([]string, 0)
	for _, e := range e.entrys {
		paths = append(paths, e.String())
	}
	return strings.Join(paths, pathListSeparator)
}