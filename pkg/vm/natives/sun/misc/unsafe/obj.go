package unsafe

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

func getInt(f *stack.Frame) {
	o := f.GetLocalRef(1)
	offset := int(f.GetLocalLong(2))

	if o.Class.IsArray() {
		f.PushOpstackVal(o.Ints()[offset])
	} else {
		f.PushOpstackVal(o.Fields().GetVal(offset))
	}
}

func init()  {
	natives.Register("sun/misc/Unsafe", "getIntVolatile", "(Ljava/lang/Object;J)I", getInt)
	natives.Register("sun/misc/Unsafe", "getInt", "(Ljava/lang/Object;J)I", getInt)
}