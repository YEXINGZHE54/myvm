package stack

import (
	"myvm/pkg/vm/engine/data"
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
}

func (f *Frame) SetLocalVal(idx int, val int) {
	f.localVars[idx].val = val
}

func (f *Frame) GetLocalVal(idx int) int {
	return f.localVars[idx].val
}

func (f *Frame) SetLocalRef(idx int, ref *data.Object) {
	f.localVars[idx].ref = ref
}

func (f *Frame) GetLocalRef(idx int) *data.Object {
	return f.localVars[idx].ref
}

func (f *Frame) PushOpstackVal(val int) {
	f.opStack.slots[f.opStack.top].val = val
	f.opStack.top = f.opStack.top + 1
}

func (f *Frame) PopOpstackVal() int {
	f.opStack.top = f.opStack.top - 1
	return f.opStack.slots[f.opStack.top].val
}

func (f *Frame) PushOpstackRef(ref *data.Object) {
	f.opStack.slots[f.opStack.top].ref = ref
	f.opStack.top = f.opStack.top + 1
}

func (f *Frame) PopOpstackRef() *data.Object {
	f.opStack.top = f.opStack.top - 1
	return f.opStack.slots[f.opStack.top].ref
}