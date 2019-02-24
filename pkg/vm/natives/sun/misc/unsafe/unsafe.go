package unsafe

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
	"unsafe"
)

func addressSize(f *stack.Frame)  {
	this := f.GetLocalRef(0)
	utils.Log("address size: %d", this)
	f.PushOpstackVal(int32(unsafe.Sizeof(this)))
}

func init()  {
	natives.Register("sun/misc/Unsafe", "addressSize", "()I", addressSize)
}