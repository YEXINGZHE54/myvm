package class

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

func getPrimitiveClass(f *stack.Frame)  {
	clsname := f.GetLocalRef(0).GoString()
	cls, err := f.GetMethod().Cls.Loader.LoadClass(clsname)
	if err != nil {
		panic(err)
	}
	f.PushOpstackRef(cls.ToObject())
}

func init()  {
	natives.Register("java/lang/Class", "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
}