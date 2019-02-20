package stack

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
)

func (f *Frame) SetPC(pc int) {
	f.pc = pc
}

func (f *Frame) GetPC() int {
	return f.pc
}

func (f *Frame) SetLocalVal(idx int, val int32) {
	f.localVars.SetVal(idx, val)
}

func (f *Frame) GetLocalVal(idx int) int32 {
	return f.localVars.GetVal(idx)
}

func (f *Frame) SetLocalRef(idx int, ref *reflect.Object) {
	f.localVars.SetRef(idx, ref)
}

func (f *Frame) GetLocalRef(idx int) *reflect.Object {
	return f.localVars.GetRef(idx)
}

func (f *Frame) SetLocalSlot(idx int, s reflect.Slot) {
	f.localVars[idx] = s
}

func (f *Frame) PushOpstackVal(val int32) {
	f.opStack.slots.SetVal(f.opStack.top, val)
	f.opStack.top = f.opStack.top + 1
}

func (f *Frame) PopOpstackVal() int32 {
	f.opStack.top = f.opStack.top - 1
	return f.opStack.slots.GetVal(f.opStack.top)
}

func (f *Frame) PushOpstackLong(val int64) {
	f.opStack.slots.SetLong(f.opStack.top, val)
	f.opStack.top = f.opStack.top + 2
}

func (f *Frame) PopOpstackLong() int64 {
	f.opStack.top = f.opStack.top - 2
	return f.opStack.slots.GetLong(f.opStack.top)
}

func (f *Frame) PushOpstackFloat(val float32) {
	f.opStack.slots.SetFloat(f.opStack.top, val)
	f.opStack.top = f.opStack.top + 1
}

func (f *Frame) PopOpstackFloat() float32 {
	f.opStack.top = f.opStack.top - 1
	return f.opStack.slots.GetFloat(f.opStack.top)
}

func (f *Frame) PushOpstackDouble(val float64) {
	f.opStack.slots.SetDouble(f.opStack.top, val)
	f.opStack.top = f.opStack.top + 2
}

func (f *Frame) PopOpstackDouble() float64 {
	f.opStack.top = f.opStack.top - 2
	return f.opStack.slots.GetDouble(f.opStack.top)
}

func (f *Frame) PushOpstackRef(ref *reflect.Object) {
	f.opStack.slots.SetRef(f.opStack.top, ref)
	f.opStack.top = f.opStack.top + 1
}

func (f *Frame) PopOpstackRef() *reflect.Object {
	f.opStack.top = f.opStack.top - 1
	return f.opStack.slots.GetRef(f.opStack.top)
}

func (f *Frame) PushOpstackSlot(s reflect.Slot) {
	f.opStack.slots.SetVal(f.opStack.top, s.Val)
	f.opStack.slots.SetRef(f.opStack.top, s.Ref)
	f.opStack.top = f.opStack.top + 1
}

func (f *Frame) DupStack() {
	idx := f.opStack.top
	f.opStack.slots[idx] = f.opStack.slots[idx-1]
	f.opStack.top = idx + 1
}

func (f *Frame) PopOpstackSlot() reflect.Slot {
	f.opStack.top = f.opStack.top - 1
	return f.opStack.slots[f.opStack.top]
}

func (f *Frame) GetOpstackSlot(idx int) reflect.Slot {
	return f.opStack.slots[f.opStack.top-1-idx]
}

func (f *Frame) ClearOpstack() {
	f.opStack.top = 0
}