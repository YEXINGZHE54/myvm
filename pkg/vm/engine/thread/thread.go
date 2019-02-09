package thread

import (
	myvm "myvm/pkg/vm"
	"myvm/pkg/vm/memory/stack"
)

func NewThread(max int, vm myvm.VM) *Thread {
	return &Thread{
		vm: vm,
		stack: stack.NewStack(max),
	}
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