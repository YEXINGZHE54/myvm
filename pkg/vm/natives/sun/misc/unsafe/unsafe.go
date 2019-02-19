package unsafe

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

func addressSize(f *stack.Frame)  {
	f.PushOpstackVal(1)
}

func init()  {
	natives.Register("sun/misc/Unsafe", "addressSize", "()I", addressSize)
}