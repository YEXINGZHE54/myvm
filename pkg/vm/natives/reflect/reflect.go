package reflect

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

func getCallerClass(f *stack.Frame)  {
	f.PushOpstackRef(f.Stack.Caller().GetMethod().Cls.ToObject())
}

func init()  {
	natives.Register("sun/reflect/Reflection", "getCallerClass", "()Ljava/lang/Class;", getCallerClass)
}