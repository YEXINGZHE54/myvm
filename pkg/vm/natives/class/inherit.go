package class

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

func isAssignableFrom(f *stack.Frame)  {
	this := f.This().Extra.(*reflect.Class)
	from := f.GetLocalRef(1).Extra.(*reflect.Class)
	if reflect.CanCastTo(from, this) {
		f.PushOpstackVal(1)
	} else {
		f.PushOpstackVal(0)
	}
}

func isInterface(f *stack.Frame)  {
	this := f.This().Extra.(*reflect.Class)
	if this.IsInterface() {
		f.PushOpstackVal(1)
	} else {
		f.PushOpstackVal(0)
	}
}

func init()  {
	natives.Register("java/lang/Class", "isAssignableFrom", "(Ljava/lang/Class;)Z", isAssignableFrom)
	natives.Register("java/lang/Class", "isInterface", "()Z", isInterface)
}