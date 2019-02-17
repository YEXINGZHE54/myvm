package convert

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	i2l_op = 0x85
)

type (
	i2lInst struct {
	}
)

func (i *i2lInst) Clone() instructions.Inst {
	return i
}

func (i *i2lInst) Fetch(coder *instructions.CodeReader) {
	return
}

func (i *i2lInst) Exec(f *stack.Frame)  {
	ival := f.PopOpstackVal()
	f.PushOpstackLong(int64(ival))
}

func init()  {
	instructions.Register(i2l_op, &i2lInst{})
}