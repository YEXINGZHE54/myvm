package constants

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	opcode_noop = 0x00
)

type (
	NoopInst struct{}
)

func (i *NoopInst) Clone() instructions.Inst {
	return i
}

func (i *NoopInst) Fetch(coder *instructions.CodeReader) {

}

func (i *NoopInst) Exec(f *stack.Frame) {

}

func init() {
	instructions.Register(opcode_noop, &NoopInst{})
}