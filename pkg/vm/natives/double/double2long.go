package double

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
	"math"
)

func doubleToRawLongBits(f *stack.Frame)  {
	v := int64(math.Float64bits(float64(f.GetLocalVal(0))))
	f.PushOpstackLong(v)
}

func init()  {
	natives.Register("java/lang/Double", "doubleToRawLongBits", "(D)J", doubleToRawLongBits)
}
