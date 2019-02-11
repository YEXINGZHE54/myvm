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

func (f *Frame) Exited() bool {
	return f.pc < 0
}

func (f *Frame) Exit() {
	f.pc = -1
	f.stack.Pop()
}

func (f *Frame) SetLocalVal(idx uint, val int32) {
	f.localVars.SetVal(idx, val)
}

func (f *Frame) GetLocalVal(idx uint) int32 {
	return f.localVars.GetVal(idx)
}

func (f *Frame) SetLocalRef(idx uint, ref *reflect.Object) {
	f.localVars.SetRef(idx, ref)
}

func (f *Frame) GetLocalRef(idx uint) *reflect.Object {
	return f.localVars.GetRef(idx)
}

func (f *Frame) PushOpstackVal(val int32) {
	f.opStack.slots.SetVal(f.opStack.top, val)
	f.opStack.top = f.opStack.top + 1
}

func (f *Frame) PopOpstackVal() int32 {
	f.opStack.top = f.opStack.top - 1
	return f.opStack.slots.GetVal(f.opStack.top)
}

func (f *Frame) PushOpstackRef(ref *reflect.Object) {
	f.opStack.slots.SetRef(f.opStack.top, ref)
	f.opStack.top = f.opStack.top + 1
}

func (f *Frame) PopOpstackRef() *reflect.Object {
	f.opStack.top = f.opStack.top - 1
	return f.opStack.slots.GetRef(f.opStack.top)
}

func (f *Frame) DupStack() {
	idx := f.opStack.top
	f.opStack.slots[idx] = f.opStack.slots[idx-1]
	f.opStack.top = idx + 1
}