package math

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	land_op = 0x7f
	iand_op = 0x7e
)

type (
	landInst struct {
	}
	iandInst struct {
	}
)

func (i *landInst) Clone() instructions.Inst {
	return i
}

func (i *landInst) Fetch(coder *instructions.CodeReader) {

}

func (i *landInst) Exec(f *stack.Frame) {
	v2 := f.PopOpstackLong()
	v1 := f.PopOpstackLong()
	f.PushOpstackLong(v1 & v2)
}

func (i *iandInst) Clone() instructions.Inst {
	return i
}

func (i *iandInst) Fetch(coder *instructions.CodeReader) {

}

func (i *iandInst) Exec(f *stack.Frame) {
	v2 := f.PopOpstackVal()
	v1 := f.PopOpstackVal()
	f.PushOpstackVal(v1 & v2)
}

func init()  {
	instructions.Register(land_op, &landInst{})
	instructions.Register(iand_op, &iandInst{})
}