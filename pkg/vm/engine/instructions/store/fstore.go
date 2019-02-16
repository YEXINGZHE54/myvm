package store

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	fstore0_op = 0x43
	fstore1_op = 0x44
	fstore2_op = 0x45
	fstore3_op = 0x46
)

type (
	FStore0 struct{
	}
	FStore1 struct{
	}
	FStore2 struct{
	}
	FStore3 struct{
	}
)

func (i *FStore0) Clone() instructions.Inst {
	return i
}

func (i *FStore0) Fetch(coder *instructions.CodeReader) {

}

func (i *FStore0) Exec(f *stack.Frame) {
	fstore(f, 0)
}

func (i *FStore1) Clone() instructions.Inst {
	return i
}

func (i *FStore1) Fetch(coder *instructions.CodeReader) {

}

func (i *FStore1) Exec(f *stack.Frame) {
	fstore(f, 1)
}

func (i *FStore2) Clone() instructions.Inst {
	return i
}

func (i *FStore2) Fetch(coder *instructions.CodeReader) {

}

func (i *FStore2) Exec(f *stack.Frame) {
	fstore(f, 2)
}

func (i *FStore3) Clone() instructions.Inst {
	return i
}

func (i *FStore3) Fetch(coder *instructions.CodeReader) {

}

func (i *FStore3) Exec(f *stack.Frame) {
	fstore(f, 3)
}

func fstore(f *stack.Frame, idx int)  {
	val := f.PopOpstackFloat()
	f.SetLocalVal(idx, int32(val))
}

func init() {
	instructions.Register(fstore0_op, &FStore0{})
	instructions.Register(fstore1_op, &FStore1{})
	instructions.Register(fstore2_op, &FStore2{})
	instructions.Register(fstore3_op, &FStore3{})
}
