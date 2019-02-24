package unsafe

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
	"sync/atomic"
	"unsafe"
)

func compareAndSwapObject(f *stack.Frame)  {
	obj := f.GetLocalRef(1)
	offset := f.GetLocalLong(2)
	expected := f.GetLocalRef(4)
	updated := f.GetLocalRef(5)
	var success bool
	if obj.Class.IsArray() {
		objs := obj.Refs()
		ps := *((*[]unsafe.Pointer)(unsafe.Pointer(&objs)))
		success = atomic.CompareAndSwapPointer(&ps[offset], unsafe.Pointer(expected), unsafe.Pointer(updated))
	} else {
		fields := obj.Fields()
		if reflect.GetRef(fields[offset]) == expected {
			fields[offset] = updated
			success = true
		}
	}
	if success {
		f.PushOpstackVal(1)
	} else {
		f.PushOpstackVal(0)
	}
}

func compareAndSwapInt(f *stack.Frame)  {
	obj := f.GetLocalRef(1)
	offset := f.GetLocalLong(2)
	expected := f.GetLocalVal(4)
	updated := f.GetLocalVal(5)
	var success bool
	if obj.Class.IsArray() {
		ints := obj.Ints()
		success = atomic.CompareAndSwapInt32(&ints[offset], expected, updated)
	} else {
		fields := obj.Fields()
		v := reflect.GetVal(fields[offset])
		success = atomic.CompareAndSwapInt32(&v, expected, updated)
	}
	if success {
		f.PushOpstackVal(1)
	} else {
		f.PushOpstackVal(0)
	}
}

func init()  {
	natives.Register("sun/misc/Unsafe", "compareAndSwapObject", "(Ljava/lang/Object;JLjava/lang/Object;Ljava/lang/Object;)Z", compareAndSwapObject)
	natives.Register("sun/misc/Unsafe", "compareAndSwapInt", "(Ljava/lang/Object;JII)Z", compareAndSwapInt)
}