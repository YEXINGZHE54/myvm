package thread

import (
	"fmt"
	myvm "github.com/YEXINGZHE54/myvm/pkg/vm"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

func NewThread(max int, vm myvm.VM, class string) *Thread {
	t := &Thread{
		vm: vm,
		class: class,
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
		fmt.Printf(">> pc:%4d %v.%v%v \n", f.GetPC(), m.Cls.Name, m.Name, m.Desc)
	}
}