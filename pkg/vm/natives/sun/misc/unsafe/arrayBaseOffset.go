package unsafe

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

func arrayBaseOffset(f *stack.Frame)  {
	f.PushOpstackVal(1)
}

func arrayIndexScale(f *stack.Frame)  {
	f.PushOpstackVal(1)
}

func init()  {
	natives.Register("sun/misc/Unsafe", "arrayBaseOffset", "(Ljava/lang/Class;)I", arrayBaseOffset)
	natives.Register("sun/misc/Unsafe", "arrayIndexScale", "(Ljava/lang/Class;)I", arrayIndexScale)
}