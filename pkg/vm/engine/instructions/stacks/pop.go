package stacks

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	pop_op = 0x57
	pop2_op = 0x58
)

type (
	PopInst struct {
	}
	Pop2Inst struct {
	}
)

func (i *PopInst) Clone() instructions.Inst {
	return i
}

func (i *PopInst) Fetch(coder *instructions.CodeReader) {

}

func (i *PopInst) Exec(f *stack.Frame) {
	println("pop exec")
	f.PopOpstackVal()
}

func (i *Pop2Inst) Clone() instructions.Inst {
	return i
}

func (i *Pop2Inst) Fetch(coder *instructions.CodeReader) {

}

func (i *Pop2Inst) Exec(f *stack.Frame) {
	f.PopOpstackVal()
	f.PopOpstackVal()
}

func init()  {
	instructions.Register(pop_op, &PopInst{})
	instructions.Register(pop2_op, &Pop2Inst{})
}