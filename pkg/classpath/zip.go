package classpath

import (
	"path/filepath"
	"archive/zip"
	"io/ioutil"
)

type ZipEntry struct {
	path string
}

func newZipEntry(path string) *ZipEntry {
	e := new(ZipEntry)
	abspath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	e.path = abspath
	return e
}

func (e *ZipEntry) readClass(clsname string) (buf []byte, ety Entry, err error) {
	r, err := zip.OpenReader(e.path)
	if err != nil {
		return
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == clsname {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			buf, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return buf, e, nil
		}
	}
	return nil, nil, ErrorClassNotFound
}

func (e *ZipEntry) String() string {
	return e.path
}