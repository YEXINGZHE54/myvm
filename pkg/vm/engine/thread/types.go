package thread

import (
	"myvm/pkg/vm/memory/stack"
	myvm "myvm/pkg/vm"
)

type (
	Thread struct {
		pc int
		vm myvm.VM
		stack *stack.Stack
	}
)