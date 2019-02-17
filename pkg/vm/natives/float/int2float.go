package float

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
	"math"
)

func intBitsToFloat(f *stack.Frame)  {
	v := int32(math.Float32frombits(uint32(f.GetLocalVal(0))))
	f.PushOpstackVal(v)
}

func init()  {
	natives.Register("java/lang/Float", "intBitsToFloat", "(I)F", intBitsToFloat)
}
