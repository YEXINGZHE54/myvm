package stacks

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	dup_op = 0x59
	dupx1_op = 0x5a
)

type (
	DupInst struct {
	}
	Dupx1Inst struct {
	}
)

func (i *DupInst) Clone() instructions.Inst {
	return i
}

func (i *DupInst) Fetch(coder *instructions.CodeReader) {

}

func (i *DupInst) Exec(f *stack.Frame) {
	println("dup op exec")
	f.DupStack()
}

func (i *Dupx1Inst) Clone() instructions.Inst {
	return i
}

func (i *Dupx1Inst) Fetch(coder *instructions.CodeReader) {

}

func (i *Dupx1Inst) Exec(f *stack.Frame) {
	println("dupx1 op exec")
	v1 := f.PopOpstackSlot()
	v2 := f.PopOpstackSlot()
	f.PushOpstackSlot(v1)
	f.PushOpstackSlot(v2)
	f.PushOpstackSlot(v1)
}

func init() {
	instructions.Register(dup_op, &DupInst{})
	instructions.Register(dupx1_op, &Dupx1Inst{})
}
