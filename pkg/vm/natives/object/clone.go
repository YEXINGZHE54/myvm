package object

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

func clone(f *stack.Frame)  {
	this := f.GetLocalRef(0)
	f.PushOpstackRef(this.Clone())
}

func init()  {
	natives.Register("java/lang/Object", "clone", "()Ljava/lang/Object;", clone)
}