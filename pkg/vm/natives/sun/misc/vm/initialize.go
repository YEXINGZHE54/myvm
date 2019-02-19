package vm

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

func initialize(f *stack.Frame)  {
	cls := f.GetMethod().Cls
	field, err := cls.LookupStaticField("savedProps", "Ljava/util/Properties;")
	if err != nil {
		panic(err)
	}
	v := cls.GetField(field).(*reflect.Object)
	// set prop
	key, err := cls.Loader.JString("for")
	if err != nil {
		panic(err)
	}
	val, err := cls.Loader.JString("bar")
	if err != nil {
		panic(err)
	}
	f.PushOpstackRef(v)
	f.PushOpstackRef(key)
	f.PushOpstackRef(val)
	propcls, err := cls.Loader.LoadClass("java/util/Properties")
	if err != nil {
		panic(err)
	}
	method, err := propcls.LookupMethod("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	if err != nil {
		panic(err)
	}
	newf := stack.NewFrame(method)
	f.Stack.Push(newf)
	newf.SetLocalSlot(2, f.PopOpstackSlot())
	newf.SetLocalSlot(1, f.PopOpstackSlot())
	newf.SetLocalSlot(0, f.PopOpstackSlot())
}

func init()  {
	natives.Register("sun/misc/VM", "initialize", "()V", initialize)
}