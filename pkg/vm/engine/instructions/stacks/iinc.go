package stacks

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	iinc_op = 0x84
)

type (
	IincInst struct {
		idx uint8
		cst int8
	}
)

func (i *IincInst) Clone() instructions.Inst {
	return &IincInst{}
}

func (i *IincInst) Fetch(coder *instructions.CodeReader) {
	i.idx = coder.Read1()
	i.cst = int8(coder.Read1())
}

func (i *IincInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction iinc")
	v := f.GetLocalVal(int(i.idx))
	f.SetLocalVal(int(i.idx), v + int32(i.cst))
}

func init()  {
	instructions.Register(iinc_op, &IincInst{})
}