package objects

import (
	"myvm/pkg/vm/memory/stack"
	"myvm/pkg/vm/loader/classfile"
	"myvm/pkg/vm/engine/instructions"
)

const (
	invokevirtual_op = 0xb6
)

type (
	InvokeVirtualInst struct {
		idx uint16
	}
)

func (i *InvokeVirtualInst) Clone() instructions.Inst {
	return &InvokeVirtualInst{}
}

func (i *InvokeVirtualInst) Fetch(coder *classfile.CodeReader) {
	i.idx = coder.Read2()
}

func (i *InvokeVirtualInst) Exec(f *stack.Frame) {
	println("invoke virtual exec")
	println(i.idx)
	//f.PushOpstackRef(nil)
}

func init() {
	instructions.Register(invokevirtual_op, &InvokeVirtualInst{})
}