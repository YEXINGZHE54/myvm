package thread

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	myvm "github.com/YEXINGZHE54/myvm/pkg/vm"
)

type (
	Thread struct {
		pc int
		boot reflect.Loader
		vm myvm.VM
		stack *stack.Stack
		class string
		args []string
	}
)