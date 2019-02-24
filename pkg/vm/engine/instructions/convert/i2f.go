package convert

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	i2f_op = 0x86
)

type (
	i2fInst struct {}
)

func (i *i2fInst) Clone() instructions.Inst {
	return i
}

func (i *i2fInst) Fetch(coder *instructions.CodeReader) {
	return
}

func (i *i2fInst) Exec(f *stack.Frame)  {
	utils.Log("executing instruction i2f")
	ival := f.PopOpstackVal()
	f.PushOpstackFloat(float32(ival))
}

func init()  {
	instructions.Register(i2f_op, &i2fInst{})
}