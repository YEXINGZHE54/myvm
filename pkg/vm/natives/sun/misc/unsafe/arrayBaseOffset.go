package unsafe

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

func arrayBaseOffset(f *stack.Frame)  {
	f.PushOpstackVal(0)
}

func arrayIndexScale(f *stack.Frame)  {
	f.PushOpstackVal(1)
}

func objectFieldOffset(f *stack.Frame) {
	jField := f.GetLocalRef(1)

	field, err := jField.Class.LookupInstanceField("slot", "I")
	if err != nil {
		panic(err)
	}
	offset := jField.GetField(field).(int32)

	f.PushOpstackLong(int64(offset))
}

func init()  {
	natives.Register("sun/misc/Unsafe", "arrayBaseOffset", "(Ljava/lang/Class;)I", arrayBaseOffset)
	natives.Register("sun/misc/Unsafe", "arrayIndexScale", "(Ljava/lang/Class;)I", arrayIndexScale)
	natives.Register("sun/misc/Unsafe", "objectFieldOffset", "(Ljava/lang/reflect/Field;)J", objectFieldOffset)
}