package vm

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

func initialize(f *stack.Frame)  {
	cls, err := f.GetMethod().Cls.Loader.LoadClass("java/lang/System")
	if err != nil {
		panic(err)
	}
	method, err := cls.LookupMethod("initializeSystemClass", "()V")
	if err != nil {
		panic(err)
	}
	newf := stack.NewFrame(method)
	f.Stack.Push(newf)
}

func init()  {
	natives.Register("sun/misc/VM", "initialize", "()V", initialize)
}