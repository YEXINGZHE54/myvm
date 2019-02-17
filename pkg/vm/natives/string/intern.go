package string

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

func intern(f *stack.Frame)  {
	loader := f.GetMethod().Cls.Loader
	this := f.GetLocalRef(0)
	loader.JString(this.GoString())
}

func init()  {
	natives.Register("java/lang/String", "intern", "()Ljava/lang/String;", intern)
}