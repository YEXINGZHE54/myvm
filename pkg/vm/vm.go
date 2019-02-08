package vm

import (
	"fmt"
	"strings"
	"myvm/pkg/vm/loader/classpath"
)

type (
	VM interface {
		Startup(class string, args []string) error
	}
	VMImpl struct {
		cp *classpath.ClassPath
	}
)

func NewVM(bootPath, classPath string) VM {
	cp := classpath.ParseOption(bootPath, classPath)
	return &VMImpl{cp}
}

func (vm *VMImpl) Startup(class string, args []string) (err error) {
	clsname := strings.Replace(class, ".", "/", -1)
	data, _, err := vm.cp.ReadClass(clsname)
	if err != nil {
		panic(err)
	}
	fmt.Printf("class byte: %v\n", data)
	return
}

