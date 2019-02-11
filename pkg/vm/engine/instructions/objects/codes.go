package objects

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
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

func (i *GetStaticInst) Fetch(coder *instructions.CodeReader) {
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