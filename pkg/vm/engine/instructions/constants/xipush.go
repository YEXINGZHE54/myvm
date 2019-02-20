package constants

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	bipush_op = 0x10
	sipush_op = 0x11
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

func (i *BIPushInst) Fetch(coder *instructions.CodeReader) {
	i.val = int8(coder.Read1())
}

func (i *BIPushInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction bipush")
	f.PushOpstackVal(int32(i.val))
}

func (i *SIPushInst) Clone() instructions.Inst {
	return &SIPushInst{}
}

func (i *SIPushInst) Fetch(coder *instructions.CodeReader) {
	i.val = int16(coder.Read2())
}

func (i *SIPushInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction sipush")
	f.PushOpstackVal(int32(i.val))
}

func init() {
	instructions.Register(bipush_op, &BIPushInst{})
	instructions.Register(sipush_op, &SIPushInst{})
}