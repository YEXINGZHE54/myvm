package vm

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
)

type (
	VM interface {
		Startup(class string, args []string) error
		LoadClass(class string) (cf *reflect.Class, err error)
	}
	VMConstructor func(bootPath, classPath string) VM
)

var (
	factory VMConstructor
)

func NewVM(bootPath, classPath string) VM {
	return factory(bootPath, classPath)
}

func Register(f VMConstructor) {
	factory = f
}