package math

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	ixor_op = 0x82
)

type (
	IXorInst struct {

	}
)

func (i *IXorInst) Clone() instructions.Inst {
	return i
}

func (i *IXorInst) Fetch(coder *instructions.CodeReader) {

}

func (i *IXorInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction xor")
	v2 := f.PopOpstackVal()
	v1 := f.PopOpstackVal()
	f.PushOpstackVal(v1 ^ v2)
}

func init()  {
	instructions.Register(ixor_op, &IXorInst{})
}