package float

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
	"math"
)

func floatToRawIntBits(f *stack.Frame)  {
	v := int32(math.Float32bits(f.GetLocalFloat(0)))
	f.PushOpstackVal(v)
}

func init()  {
	natives.Register("java/lang/Float", "floatToRawIntBits", "(F)I", floatToRawIntBits)
}
