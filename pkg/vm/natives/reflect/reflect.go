package reflect

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

func getCallerClass(f *stack.Frame)  {
	f.PushOpstackRef(f.Stack.Caller().GetMethod().Cls.ToObject())
}

// static method
func getClassAccessFlags(f *stack.Frame)  {
	cls := f.This().Extra.(*reflect.Class)
	f.PushOpstackVal(int32(cls.Flag))
}

func init()  {
	natives.Register("sun/reflect/Reflection", "getCallerClass", "()Ljava/lang/Class;", getCallerClass)
	natives.Register("sun/reflect/Reflection", "getClassAccessFlags", "(Ljava/lang/Class;)I", getClassAccessFlags)
}