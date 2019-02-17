package math

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	land_op = 0x7f
)

type (
	landInst struct {
	}
)

func (i *landInst) Clone() instructions.Inst {
	return i
}

func (i *landInst) Fetch(coder *instructions.CodeReader) {

}

func (i *landInst) Exec(f *stack.Frame) {
	v2 := f.PopOpstackLong()
	v1 := f.PopOpstackLong()
	f.PushOpstackLong(v1 & v2)
}

func init()  {
	instructions.Register(land_op, &landInst{})
}