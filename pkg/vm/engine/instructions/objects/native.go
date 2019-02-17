package objects

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

const (
	invokenative_op = 0xfe
)

type (
	NativeInst struct {
	}
)

func (i *NativeInst) Clone() instructions.Inst {
	return &NativeInst{}
}

func (i *NativeInst) Fetch(coder *instructions.CodeReader) {
}

func (i *NativeInst) Exec(f *stack.Frame) {
	m := f.GetMethod()
	nativem, err := natives.LookUpNative(m.Cls.Name, m.Name, m.Desc)
	if err != nil {
		panic(err)
	}
	nativem(f)
}

func init()  {
	instructions.Register(invokenative_op, &NativeInst{})
}