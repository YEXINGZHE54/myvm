package constants

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/loader/classfile"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
)

const (

)

type (
	BIPushInst struct{
		val int8
	}
	SIPushInst struct{
		val int16
	}
)

func (i *BIPushInst) Clone() instructions.Inst {
	return &BIPushInst{}
}

func (i *BIPushInst) Fetch(coder *classfile.CodeReader) {
	i.val = int8(coder.Read1())
}

func (i *BIPushInst) Exec(f *stack.Frame) {
	f.PushOpstackVal(int(i.val))
}

func (i *SIPushInst) Clone() instructions.Inst {
	return &SIPushInst{}
}

func (i *SIPushInst) Fetch(coder *classfile.CodeReader) {
	i.val = int16(coder.Read2())
}

func (i *SIPushInst) Exec(f *stack.Frame) {
	f.PushOpstackVal(int(i.val))
}

func init() {
	instructions.Register(opcode_noop, &BIPushInst{})
	instructions.Register(opcode_noop, &SIPushInst{})
}