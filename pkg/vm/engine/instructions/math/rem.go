package math

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	irem_op = 0x70
)

type (
	IRemInst struct {

	}
)

func (i *IRemInst) Clone() instructions.Inst {
	return i
}

func (i *IRemInst) Fetch(coder *instructions.CodeReader) {

}

func (i *IRemInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction irem")
	v2 := f.PopOpstackVal()
	v1 := f.PopOpstackVal()
	f.PushOpstackVal(v1 % v2)
}

func init()  {
	instructions.Register(irem_op, &IRemInst{})
}