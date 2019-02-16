package constants

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	constnull_op = 0x01
	dconst0_op = 0x0e
	dconst1_op = 0x0f
	fconst0_op = 0x0b
	fconst1_op = 0x0c
	fconst2_op = 0x0d
	iconstm1_op = 0x02
	iconst0_op = 0x03
	iconst1_op = 0x04
	iconst2_op = 0x05
	iconst3_op = 0x06
	iconst4_op = 0x07
	iconst5_op = 0x08
	lcosnt0_op = 0x09
	lconst1_op = 0x0a
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

func (i *ConstNullInst) Fetch(coder *instructions.CodeReader) {

}

func (i *ConstNullInst) Exec(f *stack.Frame) {
	f.PushOpstackRef(nil)
}

func (i *DConst0Inst) Clone() instructions.Inst {
	return i
}

func (i *DConst0Inst) Fetch(coder *instructions.CodeReader) {

}

func (i *DConst0Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(0)
	f.PushOpstackVal(0) // just a slot
}

func (i *DConst1Inst) Clone() instructions.Inst {
	return i
}

func (i *DConst1Inst) Fetch(coder *instructions.CodeReader) {

}

func (i *DConst1Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(1)
	f.PushOpstackVal(1) // just a alot
}

func (i *FConst0Inst) Clone() instructions.Inst {
	return i
}

func (i *FConst0Inst) Fetch(coder *instructions.CodeReader) {

}

func (i *FConst0Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(0)
}

func (i *FConst1Inst) Clone() instructions.Inst {
	return i
}

func (i *FConst1Inst) Fetch(coder *instructions.CodeReader) {

}

func (i *FConst1Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(1)
}

func (i *FConst2Inst) Clone() instructions.Inst {
	return i
}

func (i *FConst2Inst) Fetch(coder *instructions.CodeReader) {

}

func (i *FConst2Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(2)
}

func (i *IConstM1Inst) Clone() instructions.Inst {
	return i
}

func (i *IConstM1Inst) Fetch(coder *instructions.CodeReader) {

}

func (i *IConstM1Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(-1)
}

func (i *IConst0Inst) Clone() instructions.Inst {
	return i
}

func (i *IConst0Inst) Fetch(coder *instructions.CodeReader) {

}

func (i *IConst0Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(0)
}

func (i *IConst1Inst) Clone() instructions.Inst {
	return i
}

func (i *IConst1Inst) Fetch(coder *instructions.CodeReader) {

}

func (i *IConst1Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(1)
}

func (i *IConst2Inst) Clone() instructions.Inst {
	return i
}

func (i *IConst2Inst) Fetch(coder *instructions.CodeReader) {

}

func (i *IConst2Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(2)
}

func (i *IConst3Inst) Clone() instructions.Inst {
	return i
}

func (i *IConst3Inst) Fetch(coder *instructions.CodeReader) {

}

func (i *IConst3Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(3)
}

func (i *IConst4Inst) Clone() instructions.Inst {
	return i
}

func (i *IConst4Inst) Fetch(coder *instructions.CodeReader) {

}

func (i *IConst4Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(4)
}

func (i *IConst5Inst) Clone() instructions.Inst {
	return i
}

func (i *IConst5Inst) Fetch(coder *instructions.CodeReader) {

}

func (i *IConst5Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(5)
}

func (i *LConst0Inst) Clone() instructions.Inst {
	return i
}

func (i *LConst0Inst) Fetch(coder *instructions.CodeReader) {

}

func (i *LConst0Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(0)
	f.PushOpstackVal(0) // just slot
}

func (i *LConst1Inst) Clone() instructions.Inst {
	return i
}

func (i *LConst1Inst) Fetch(coder *instructions.CodeReader) {

}

func (i *LConst1Inst) Exec(f *stack.Frame) {
	f.PushOpstackVal(1)
	f.PushOpstackVal(1) // just slot
}

func init() {
	instructions.Register(constnull_op, &ConstNullInst{})
	instructions.Register(dconst0_op, &DConst0Inst{})
	instructions.Register(dconst1_op, &DConst1Inst{})
	instructions.Register(fconst0_op, &FConst0Inst{})
	instructions.Register(fconst1_op, &FConst1Inst{})
	instructions.Register(fconst2_op, &FConst2Inst{})
	instructions.Register(iconstm1_op, &IConstM1Inst{})
	instructions.Register(iconst0_op, &IConst0Inst{})
	instructions.Register(iconst1_op, &IConst1Inst{})
	instructions.Register(iconst2_op, &IConst2Inst{})
	instructions.Register(iconst3_op, &IConst3Inst{})
	instructions.Register(iconst4_op, &IConst4Inst{})
	instructions.Register(iconst5_op, &IConst5Inst{})
	instructions.Register(lcosnt0_op, &LConst0Inst{})
	instructions.Register(lconst1_op, &LConst1Inst{})
}