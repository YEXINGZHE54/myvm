package convert

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	f2i_op = 0x8b
)

type (
	f2iInst struct {}
)

func (i *f2iInst) Clone() instructions.Inst {
	return i
}

func (i *f2iInst) Fetch(coder *instructions.CodeReader) {
	return
}

func (i *f2iInst) Exec(f *stack.Frame)  {
	utils.Log("executing instruction f2i")
	ival := f.PopOpstackFloat()
	f.PushOpstackVal(int32(ival))
}

func init()  {
	instructions.Register(f2i_op, &f2iInst{})
}