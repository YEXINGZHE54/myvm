package stacks

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	anull_op = 0x01
)

type (
	ANullInst struct {
	}
)

func (i *ANullInst) Clone() instructions.Inst {
	return i
}

func (i *ANullInst) Fetch(coder *instructions.CodeReader) {

}

func (i *ANullInst) Exec(f *stack.Frame) {
	println("aconst_null op exec")
	f.PushOpstackRef(nil)
}

func init() {
	instructions.Register(dup_op, &DupInst{})
	instructions.Register(anull_op, &ANullInst{})
}
