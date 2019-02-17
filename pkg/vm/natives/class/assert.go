package class

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

func desiredAssertionStatus0(f *stack.Frame)  {
	f.PushOpstackVal(0) //false
}

func init()  {
	natives.Register("java/lang/Class", "desiredAssertionStatus0", "(Ljava/lang/Class;)Z", desiredAssertionStatus0)
}