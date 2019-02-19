package object

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
	"unsafe"
)

func hashCode(f *stack.Frame)  {
	this := f.GetLocalRef(0)
	f.PushOpstackVal(int32(uintptr(unsafe.Pointer(this))))
}

func init()  {
	natives.Register("java/lang/Object", "hashCode", "()I", hashCode)
}