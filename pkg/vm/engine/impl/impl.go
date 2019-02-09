package impl

import (
	"strings"
	"myvm/pkg/vm"
	"myvm/pkg/vm/loader/classpath"
	"github.com/kr/pretty"
	"myvm/pkg/vm/loader/classfile"
	"myvm/pkg/vm/engine/thread"
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
	clsname := strings.Replace(class, ".", "/", -1)
	data, _, err := vm.cp.ReadClass(clsname)
	if err != nil {
		return
	}
	cf, err := classfile.Parse(data)
	if err != nil {
		return
	}
	pretty.Println(cf)
	t := thread.NewThread(1024, vm)
	return t.Run()
}

func init() {
	vm.Register(NewVM)
}