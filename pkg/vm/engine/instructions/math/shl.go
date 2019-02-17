package math

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	lshl_op = 0x79
)

type (
	lshlInst struct {
	}
)

func (i *lshlInst) Clone() instructions.Inst {
	return i
}

func (i *lshlInst) Fetch(coder *instructions.CodeReader) {

}

func (i *lshlInst) Exec(f *stack.Frame) {
	v2 := f.PopOpstackVal()
	v1 := f.PopOpstackLong()
	shift := uint(v2 & 0x3F)
	f.PushOpstackLong(v1 << shift)
}

func init()  {
	instructions.Register(lshl_op, &lshlInst{})
}