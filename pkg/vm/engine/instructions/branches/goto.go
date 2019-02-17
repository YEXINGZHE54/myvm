package branches

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/thread"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

func gotoOffset(f *stack.Frame, offset int16)  {
	t := f.Stack.Thread().(*thread.Thread)
	nextpc := t.GetPC() + int(offset)
	f.SetPC(nextpc)
}