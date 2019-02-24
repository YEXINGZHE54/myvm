package impl

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/thread"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/loader/classloader"
	// including instruments
	_ "github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions/includes"
	// natives
	_ "github.com/YEXINGZHE54/myvm/pkg/vm/natives/includes"
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
	utils.Log("staring up VM, max stack: %d, classname: %s, args: %v", 256, class, args)
	t := thread.NewThread(256, vm.ld, vm, class, args)
	return t.Run()
}

func init() {
	vm.Register(NewVM)
}