package system

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

// static
func mapLibraryName(f *stack.Frame)  {
	name := f.This()
	f.PushOpstackRef(name)
}

func init()  {
	natives.Register("java/lang/System", "mapLibraryName", "(Ljava/lang/String;)Ljava/lang/String;", mapLibraryName)
}