package thread

import (
	"fmt"
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	myvm "github.com/YEXINGZHE54/myvm/pkg/vm"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/kr/pretty"
)

func NewThread(max int, boot reflect.Loader, vm myvm.VM, class string, args []string) *Thread {
	t := &Thread{
		boot: boot,
		vm: vm,
		class: class,
		args: args,
	}
	t.stack = stack.NewStack(max, t)
	return t
}

func (t *Thread) GetPC() int {
	return t.pc
}

func (t *Thread) PC(pc int) {
	t.pc = pc
}

func (t *Thread) Push(f *stack.Frame) {
	t.stack.Push(f)
}

func (t *Thread) Pop() {
	t.stack.Pop()
}

func (t *Thread) Current() (f *stack.Frame) {
	return t.stack.Current()
}

func (t *Thread) Dump() {
	fmt.Println("Java Thread Stack Dump:")
	st := t.stack
	for st.Current() != nil {
		f := st.Pop()
		m := f.GetMethod()
		fmt.Printf(">> pc:%4d Line: %4d %v.%v%v \n", f.GetPC(), m.GetLineNumber(f.GetPC()-1), m.Cls.Name, m.Name, m.Desc)
		if utils.TracingEnabled() {
			pretty.Println("instructions: ", instructions.ReadAll(m.Codes[:f.GetPC()]))
		}
	}
}