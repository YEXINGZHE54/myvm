package class

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
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

func getSuperClass(f *stack.Frame)  {
	cls := f.This().Extra.(*reflect.Class).Super
	if cls != nil {
		f.PushOpstackRef(cls.ToObject())
	} else {
		f.PushOpstackRef(nil)
	}
}

func isPrimitive(f *stack.Frame)  {
	cls := f.This().Extra.(*reflect.Class)
	if cls.IsPrimitive() {
		f.PushOpstackVal(1)
	} else {
		f.PushOpstackVal(0)
	}
}

func init()  {
	natives.Register("java/lang/Class", "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	natives.Register("java/lang/Class", "getSuperclass", "()Ljava/lang/Class;", getSuperClass)
	natives.Register("java/lang/Class", "isPrimitive", "()Z", isPrimitive)
}