package math

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	isub_op = 0x64
	lsub_op = 0x65
)

type (
	ISubInst struct {
	}
	LSubInst struct {
	}
)

func (i *ISubInst) Clone() instructions.Inst {
	return i
}

func (i *ISubInst) Fetch(coder *instructions.CodeReader)  {

}

func (i *ISubInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction isub")
	v2 := f.PopOpstackVal()
	v1 := f.PopOpstackVal()
	f.PushOpstackVal(v1 - v2)
}

func (i *LSubInst) Clone() instructions.Inst {
	return i
}

func (i *LSubInst) Fetch(coder *instructions.CodeReader)  {

}

func (i *LSubInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction lsub")
	v2 := f.PopOpstackLong()
	v1 := f.PopOpstackLong()
	f.PushOpstackLong(v1 - v2)
}

func init()  {
	instructions.Register(isub_op, &ISubInst{})
	instructions.Register(lsub_op, &LSubInst{})
}