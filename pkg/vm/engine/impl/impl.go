package impl

import (
	"myvm/pkg/vm"
	"myvm/pkg/vm/engine/thread"
	"myvm/pkg/vm/engine/reflect"
	"myvm/pkg/vm/loader/classloader"
	// including instruments
	_ "myvm/pkg/vm/engine/instructions/constants"
	_ "myvm/pkg/vm/engine/instructions/objects"
)

type (
	VMImpl struct {
		ld reflect.Loader
	}
)

func NewVM(bootPath, classPath string) vm.VM {
	return &VMImpl{classloader.NewLoader(bootPath, classPath)}
}

func (vm *VMImpl) Startup(class string, args []string) (err error) {
	t := thread.NewThread(1024, vm, class)
	return t.Run()
}

func (vm *VMImpl) LoadClass(class string) (cf *reflect.Class, err error) {
	return vm.ld.LoadClass(class)
}

func init() {
	vm.Register(NewVM)
}