package stacks

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	dup_op = 0x59
)

type (
	DupInst struct {
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

func init() {
	instructions.Register(dup_op, &DupInst{})
}
