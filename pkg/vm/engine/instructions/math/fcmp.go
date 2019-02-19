package math

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	fcmpg_op = 0x96
	fcmpl_op = 0x95
)

type (
	FcmpInst struct {
		result int
	}
)

func (i *FcmpInst) Clone() instructions.Inst {
	return i
}

func (i *FcmpInst) Fetch(coder *instructions.CodeReader)  {

}

func (i *FcmpInst) Exec(f *stack.Frame) {
	v2 := float32(f.PopOpstackVal())
	v1 := float32(f.PopOpstackVal())
	if v1 > v2 {
		f.PushOpstackVal(1)
	} else if v1 == v2 {
		f.PushOpstackVal(0)
	} else {
		f.PushOpstackVal(-1)
	}
}

func init()  {
	instructions.Register(fcmpg_op, &FcmpInst{1})
	instructions.Register(fcmpl_op, &FcmpInst{-1})
}