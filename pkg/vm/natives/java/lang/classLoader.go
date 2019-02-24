package lang

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

// static
func findBuiltinLib(f *stack.Frame)  {
	name := f.This()
	f.PushOpstackRef(name)
}

func load(f *stack.Frame)  {

}

func init()  {
	natives.Register("java/lang/ClassLoader", "findBuiltinLib", "(Ljava/lang/String;)Ljava/lang/String;", findBuiltinLib)
	natives.Register("java/lang/ClassLoader$NativeLibrary", "load", "(Ljava/lang/String;Z)V", load)
}