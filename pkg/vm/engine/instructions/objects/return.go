package objects

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	return_op = 0xb1
)

type (
	ReturnInst struct {}
)

func (i *ReturnInst) Clone() instructions.Inst {
	return i
}

func (i *ReturnInst) Fetch(coder *instructions.CodeReader) {

}

func (i *ReturnInst) Exec(f *stack.Frame) {
	println("return exec")
	f.Exit()
}

func init() {
	instructions.Register(return_op, &ReturnInst{})
}