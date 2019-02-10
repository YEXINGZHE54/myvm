package objects

import (
	"myvm/pkg/vm/memory/stack"
	"myvm/pkg/vm/loader/classfile"
	"myvm/pkg/vm/engine/instructions"
)

const (
	getstatic_op = 0xb2
)

type (
	GetStaticInst struct {
		idx uint16
	}
)

func (i *GetStaticInst) Clone() instructions.Inst {
	return &GetStaticInst{}
}

func (i *GetStaticInst) Fetch(coder *classfile.CodeReader) {
	i.idx = coder.Read2()
}

func (i *GetStaticInst) Exec(f *stack.Frame) {
	println("getstatic exec: ")
	println(i.idx)
	//f.PushOpstackRef(nil)
}

func init() {
	instructions.Register(getstatic_op, &GetStaticInst{})
}