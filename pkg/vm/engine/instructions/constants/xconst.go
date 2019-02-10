package constants

import (
	"myvm/pkg/vm/memory/stack"
	"myvm/pkg/vm/loader/classfile"
	"myvm/pkg/vm/engine/instructions"
)

const (

)

type (
	ConstNullInst struct{}
	DConst0Inst struct{}
	DConst1Inst struct{}
	FConst0Inst struct{}
	FConst1Inst struct{}
	FConst2Inst struct{}
	IConstM1Inst struct{}
	IConst0Inst struct{}
	IConst1Inst struct{}
	IConst2Inst struct{}
	IConst3Inst struct{}
	IConst4Inst struct{}
	IConst5Inst struct{}
	LConst0Inst struct{}
	LConst1Inst struct{}
)

func (i *ConstNullInst) Clone() instructions.Inst {
	return i
}

func (i *ConstNullInst) Fetch(coder *classfile.CodeReader) {

}

func (i *ConstNullInst) Exec(f *stack.Frame) {
	f.PushOpstackRef(nil)
}

func (i *DConst0Inst) Clone() instructions.Inst {
	return i
}

func (i *DConst0Inst) Fetch(coder *classfile.CodeReader) {

}

func (i *DConst0Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(0)
	f.PushOpstackVal(0) // just a slot
}

func (i *DConst1Inst) Clone() instructions.Inst {
	return i
}

func (i *DConst1Inst) Fetch(coder *classfile.CodeReader) {

}

func (i *DConst1Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(1)
	f.PushOpstackVal(1) // just a alot
}

func (i *FConst0Inst) Clone() instructions.Inst {
	return i
}

func (i *FConst0Inst) Fetch(coder *classfile.CodeReader) {

}

func (i *FConst0Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(0)
}

func (i *FConst1Inst) Clone() instructions.Inst {
	return i
}

func (i *FConst1Inst) Fetch(coder *classfile.CodeReader) {

}

func (i *FConst1Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(1)
}

func (i *FConst2Inst) Clone() instructions.Inst {
	return i
}

func (i *FConst2Inst) Fetch(coder *classfile.CodeReader) {

}

func (i *FConst2Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(2)
}

func (i *IConstM1Inst) Clone() instructions.Inst {
	return i
}

func (i *IConstM1Inst) Fetch(coder *classfile.CodeReader) {

}

func (i *IConstM1Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(-1)
}

func (i *IConst0Inst) Clone() instructions.Inst {
	return i
}

func (i *IConst0Inst) Fetch(coder *classfile.CodeReader) {

}

func (i *IConst0Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(0)
}

func (i *IConst1Inst) Clone() instructions.Inst {
	return i
}

func (i *IConst1Inst) Fetch(coder *classfile.CodeReader) {

}

func (i *IConst1Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(1)
}

func (i *IConst2Inst) Clone() instructions.Inst {
	return i
}

func (i *IConst2Inst) Fetch(coder *classfile.CodeReader) {

}

func (i *IConst2Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(2)
}

func (i *IConst3Inst) Clone() instructions.Inst {
	return i
}

func (i *IConst3Inst) Fetch(coder *classfile.CodeReader) {

}

func (i *IConst3Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(3)
}

func (i *IConst4Inst) Clone() instructions.Inst {
	return i
}

func (i *IConst4Inst) Fetch(coder *classfile.CodeReader) {

}

func (i *IConst4Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(4)
}

func (i *IConst5Inst) Clone() instructions.Inst {
	return i
}

func (i *IConst5Inst) Fetch(coder *classfile.CodeReader) {

}

func (i *IConst5Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(5)
}

func (i *LConst0Inst) Clone() instructions.Inst {
	return i
}

func (i *LConst0Inst) Fetch(coder *classfile.CodeReader) {

}

func (i *LConst0Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(0)
	f.PushOpstackVal(0) // just slot
}

func (i *LConst1Inst) Clone() instructions.Inst {
	return i
}

func (i *LConst1Inst) Fetch(coder *classfile.CodeReader) {

}

func (i *LConst1Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(1)
	f.PushOpstackVal(1) // just slot
}

func init() {
	instructions.Register(opcode_noop, &ConstNullInst{})
	instructions.Register(opcode_noop, &DConst0Inst{})
	instructions.Register(opcode_noop, &DConst1Inst{})
	instructions.Register(opcode_noop, &FConst0Inst{})
	instructions.Register(opcode_noop, &FConst1Inst{})
	instructions.Register(opcode_noop, &FConst2Inst{})
	instructions.Register(opcode_noop, &IConstM1Inst{})
	instructions.Register(opcode_noop, &IConst0Inst{})
	instructions.Register(opcode_noop, &IConst1Inst{})
	instructions.Register(opcode_noop, &IConst2Inst{})
	instructions.Register(opcode_noop, &IConst3Inst{})
	instructions.Register(opcode_noop, &IConst4Inst{})
	instructions.Register(opcode_noop, &IConst5Inst{})
	instructions.Register(opcode_noop, &LConst0Inst{})
	instructions.Register(opcode_noop, &LConst1Inst{})
}