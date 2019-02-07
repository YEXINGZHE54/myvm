package classpath

import (
	"os"
	"strings"
	"path/filepath"
)

func newWildcardEntry(path string) Entry {
	e := new(CompositeEntry)
	base := path[:len(path)-1]
	filepath.Walk(base, func (p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && p != base {
			return filepath.SkipDir
		}
		if strings.HasSuffix(p, ".jar") ||
			strings.HasSuffix(p, ".JAR") {
				e.entrys = append(e.entrys, newZipEntry(p))
			}
		return nil
	})
	return e
}