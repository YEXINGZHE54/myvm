package double

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
	"math"
)

func longBitsToDouble(f *stack.Frame)  {
	v1 := f.GetLocalVal(0)
	f.PushOpstackDouble(math.Float64frombits(uint64(v1)))
}

func init()  {
	natives.Register("java/lang/Double", "longBitsToDouble", "(J)D", longBitsToDouble)
}
