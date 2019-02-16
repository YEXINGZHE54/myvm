package impl

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/thread"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/loader/classloader"
	// including instruments
	_ "github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions/array"
	_ "github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions/constants"
	_ "github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions/objects"
	_ "github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions/stacks"
	_ "github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions/store"
	_ "github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions/load"
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
	t := thread.NewThread(1024, vm, class, args)
	return t.Run()
}

func (vm *VMImpl) LoadClass(class string) (cf *reflect.Class, err error) {
	return vm.ld.LoadClass(class)
}

func init() {
	vm.Register(NewVM)
}