package objects

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	invokevirtual_op = 0xb6
	invokespecial_op = 0xb7
)

type (
	InvokeVirtualInst struct {
		idx uint16
	}
	InvokeSpecialInst struct {
		idx uint16
	}
)

func (i *InvokeVirtualInst) Clone() instructions.Inst {
	return &InvokeVirtualInst{}
}

func (i *InvokeVirtualInst) Fetch(coder *instructions.CodeReader) {
	i.idx = coder.Read2()
}

func (i *InvokeVirtualInst) Exec(f *stack.Frame) {
	println("invoke virtual exec")
	println(i.idx)
	//f.PushOpstackRef(nil)
}

func (i *InvokeSpecialInst) Clone() instructions.Inst {
	return &InvokeSpecialInst{}
}

func (i *InvokeSpecialInst) Fetch(coder *instructions.CodeReader) {
	i.idx = coder.Read2()
}

func (i *InvokeSpecialInst) Exec(f *stack.Frame) {
	println("invoke special exec")
	println(i.idx)
	//f.PushOpstackRef(nil)
}

func init() {
	instructions.Register(invokevirtual_op, &InvokeVirtualInst{})
	instructions.Register(invokespecial_op, &InvokeVirtualInst{})
}