package string

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

func intern(f *stack.Frame)  {
	loader := f.GetMethod().Cls.Loader
	this := f.GetLocalRef(0)
	o, err := loader.JString(this.GoString())
	if err != nil {
		panic(err)
	}
	f.PushOpstackRef(o)
}

func init()  {
	natives.Register("java/lang/String", "intern", "()Ljava/lang/String;", intern)
}