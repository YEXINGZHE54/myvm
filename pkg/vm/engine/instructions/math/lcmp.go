package math

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	lcmp_op = 0x94
)

type (
	LcmpInst struct {
	}
)

func (i *LcmpInst) Clone() instructions.Inst {
	return i
}

func (i *LcmpInst) Fetch(coder *instructions.CodeReader)  {

}

func (i *LcmpInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction lcmp")
	v2 := f.PopOpstackLong()
	v1 := f.PopOpstackLong()
	if v1 > v2 {
		f.PushOpstackVal(1)
	} else if v1 == v2 {
		f.PushOpstackVal(0)
	} else {
		f.PushOpstackVal(-1)
	}
}

func init()  {
	instructions.Register(lcmp_op, &LcmpInst{})
}