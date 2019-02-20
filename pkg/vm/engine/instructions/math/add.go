package math

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	iadd_op = 0x60
	ladd_op = 0x61
)

type (
	IAddInst struct {
	}
	LAddInst struct {
	}
)

func (i *IAddInst) Clone() instructions.Inst {
	return i
}

func (i *IAddInst) Fetch(coder *instructions.CodeReader)  {

}

func (i *IAddInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction iadd")
	v2 := f.PopOpstackVal()
	v1 := f.PopOpstackVal()
	f.PushOpstackVal(v1 + v2)
}

func (i *LAddInst) Clone() instructions.Inst {
	return i
}

func (i *LAddInst) Fetch(coder *instructions.CodeReader)  {

}

func (i *LAddInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction ladd")
	v2 := f.PopOpstackLong()
	v1 := f.PopOpstackLong()
	f.PushOpstackLong(v1 + v2)
}

func init()  {
	instructions.Register(iadd_op, &IAddInst{})
	instructions.Register(ladd_op, &LAddInst{})
}