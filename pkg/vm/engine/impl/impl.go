package impl

import (
	"strings"
	"myvm/pkg/vm"
	"myvm/pkg/vm/loader/classpath"
	"myvm/pkg/vm/loader/classfile"
	"myvm/pkg/vm/engine/thread"
	// including instruments
	_ "myvm/pkg/vm/engine/instructions/constants"
	_ "myvm/pkg/vm/engine/instructions/objects"
)

type (
	VMImpl struct {
		cp *classpath.ClassPath
	}
)

func NewVM(bootPath, classPath string) vm.VM {
	cp := classpath.ParseOption(bootPath, classPath)
	return &VMImpl{cp}
}

func (vm *VMImpl) Startup(class string, args []string) (err error) {
	t := thread.NewThread(1024, vm, class)
	return t.Run()
}

func (vm *VMImpl) LoadClass(class string) (cf *classfile.ClassFile, err error) {
	clsname := strings.Replace(class, ".", "/", -1)
	data, _, err := vm.cp.ReadClass(clsname)
	if err != nil {
		return
	}
	cf, err = classfile.Parse(data)
	if err != nil {
		return
	}
	return
}

func init() {
	vm.Register(NewVM)
}