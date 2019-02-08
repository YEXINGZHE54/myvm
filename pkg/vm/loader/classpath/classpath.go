package classpath

import (
	"os"
	"path/filepath"
)

const (
	DefaultUserCP = "."
)

type ClassPath struct {
	boot Entry
	ext Entry
	user Entry
}

func ParseOption(jre, usercp string) *ClassPath {
	if usercp == "" {
		usercp = DefaultUserCP
	}
	jre = fixedJavaHome(jre)
	cp := new(ClassPath)
	cp.boot = newEntry(filepath.Join(jre, "lib", "*"))
	cp.ext = newEntry(filepath.Join(jre, "lib", "ext", "*"))
	cp.user = newEntry(usercp)
	return cp
}

func (cp *ClassPath) ReadClass(clsname string) (buf []byte, e Entry, err error) {
	clsname = clsname + ".class"
	buf, e, err = cp.boot.readClass(clsname)
	if err == nil {
		return
	}
	buf, e, err = cp.ext.readClass(clsname)
	if err == nil {
		return
	}
	buf, e, err = cp.user.readClass(clsname)
	return
}

func (cp *ClassPath) String() string {
	return cp.user.String()
}

func fixedJavaHome(jre string) string {
	if jre != "" && exists(jre) {
		return jre
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("no jre dir found!")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}