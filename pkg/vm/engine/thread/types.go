package thread

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	myvm "github.com/YEXINGZHE54/myvm/pkg/vm"
)

type (
	Thread struct {
		pc int
		vm myvm.VM
		stack *stack.Stack
		class string
	}
)