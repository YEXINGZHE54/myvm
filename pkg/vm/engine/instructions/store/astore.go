package store

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	astore1_op = 0x4c
)

type (
	Astore1Inst struct{}
)

func (i *Astore1Inst) Clone() instructions.Inst {
	return i
}

func (i *Astore1Inst) Fetch(coder *instructions.CodeReader) {

}

func (i *Astore1Inst) Exec(f *stack.Frame) {
	astore(f, 1)
}

func astore(f *stack.Frame, idx int)  {
	ref := f.PopOpstackRef()
	f.SetLocalRef(idx, ref)
}

func init() {
	instructions.Register(astore1_op, &Astore1Inst{})
}
