package classpath

import (
	"os"
	"strings"
	"errors"
)

const pathListSeparator = string(os.PathListSeparator)

var (
	ErrorClassNotFound = errors.New("class not found")
)

type (
	Entry interface {
		readClass(clsname string) ([]byte, Entry, error)
		String() string
	}
)

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	} else if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	} else if strings.HasSuffix(path, ".jar") ||
		strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") ||
		strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}