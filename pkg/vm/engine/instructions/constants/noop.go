package constants

import (
	"myvm/pkg/vm/memory/stack"
	"myvm/pkg/vm/loader/classfile"
	"myvm/pkg/vm/engine/instructions"
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

func (i *NoopInst) Fetch(coder *classfile.CodeReader) {

}

func (i *NoopInst) Exec(f *stack.Frame) {

}

func init() {
	instructions.Register(opcode_noop, &NoopInst{})
}