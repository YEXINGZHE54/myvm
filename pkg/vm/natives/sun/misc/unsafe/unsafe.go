package unsafe

import (
	"encoding/binary"
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

func allocateMemory(f *stack.Frame)  {
	size := f.GetLocalLong(1)
	f.PushOpstackLong(alloc(size))
}

func freeMemory(f *stack.Frame)  {
	addr := f.GetLocalLong(1)
	free(addr)
}

func putByte(f *stack.Frame)  {
	buf := _memaddr(f)
	buf[0] = byte(f.GetLocalVal(3))
}

func getByte(f *stack.Frame)  {
	buf := _memaddr(f)
	f.PushOpstackVal(int32(buf[0]))
}

func putLong(f *stack.Frame)  {
	buf := _memaddr(f)
	binary.BigEndian.PutUint64(buf, uint64(f.GetLocalLong(3)))
}

func getLong(f *stack.Frame)  {
	buf := _memaddr(f)
	f.PushOpstackLong(int64(binary.BigEndian.Uint64(buf)))
}

func _memaddr(f *stack.Frame) ([]byte) {
	return at(f.GetLocalLong(1))
}

func init()  {
	natives.Register("sun/misc/Unsafe", "addressSize", "()I", addressSize)
	natives.Register("sun/misc/Unsafe", "allocateMemory", "(J)J", allocateMemory)
	natives.Register("sun/misc/Unsafe", "freeMemory", "(J)V", freeMemory)
	natives.Register("sun/misc/Unsafe", "putLong", "(JJ)V", putLong)
	natives.Register("sun/misc/Unsafe", "getLong", "(J)J", getLong)
	natives.Register("sun/misc/Unsafe", "putByte", "(JB)V", putByte)
	natives.Register("sun/misc/Unsafe", "getByte", "(J)B", getByte)
}