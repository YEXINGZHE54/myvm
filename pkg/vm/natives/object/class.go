package object

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

func getClass(f *stack.Frame)  {
	f.PushOpstackRef(f.GetLocalRef(0).Class.ToObject())
}

func init()  {
	 natives.Register("java/lang/Object", "getClass", "()Ljava/lang/Class;", getClass)
}