package math

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	isub_op = 0x64
)

type (
	ISubInst struct {
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

func init()  {
	instructions.Register(isub_op, &ISubInst{})
}