package reflect

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/thread"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

func newInstance0(f *stack.Frame) {
	constructorObj := f.This()
	argArrObj := f.GetLocalRef(1)

	var ctor *reflect.Method
	if constructorObj.Extra != nil { // it's origin
		ctor = constructorObj.Extra.(*reflect.Method)
	} else {
		// if Extra is nil, it means this obj is not created by native/class/getDeclaredConstructors0 function
		// this obj is created by copy of an Origin One, which is stored in root field of this obj
		// I call this pattern as Pattern of Copy Accessor
		field, err := constructorObj.Class.LookupInstanceField("root", "Ljava/lang/reflect/Constructor;")
		if err != nil {
			panic(err)
		}
		root := constructorObj.GetField(field).(*reflect.Object)
		ctor = root.Extra.(*reflect.Method)
	}
	cls := ctor.Cls
	if !cls.Started {
		t := f.Stack.Thread().(thread.Thread)
		f.SetPC(t.GetPC())
		t.InitClass(cls)
		return
	}

	obj, err := cls.NewObject()
	if err != nil {
		panic(err)
	}
	f.PushOpstackRef(obj)

	// call <init>
	newf := stack.NewFrame(ctor)
	f.Stack.Push(newf)
	newf.SetLocalRef(0, obj)
	if argArrObj != nil {
		for idx, arg := range argArrObj.Refs() {
			newf.SetLocalRef(idx+1, arg)
		}
	}
}

func init() {
	natives.Register("sun/reflect/NativeConstructorAccessorImpl", "newInstance0", "(Ljava/lang/reflect/Constructor;[Ljava/lang/Object;)Ljava/lang/Object;", newInstance0)
}