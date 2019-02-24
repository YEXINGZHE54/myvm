package math

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	ior_op = 0x80
)

type (
	IorInst struct {

	}
)

func (i *IorInst) Clone() instructions.Inst {
	return i
}

func (i *IorInst) Fetch(coder *instructions.CodeReader) {

}

func (i *IorInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction ior")
	v2 := f.PopOpstackVal()
	v1 := f.PopOpstackVal()
	f.PushOpstackVal(v1 | v2)
}

func init()  {
	instructions.Register(ior_op, &IorInst{})
}