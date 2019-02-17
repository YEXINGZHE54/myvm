package load

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	iload_op = 0x15
	iload0_op = 0x1a
	iload1_op = 0x1b
	iload2_op = 0x1c
	iload3_op = 0x1d
)

type (
	ILoad struct {
		idx int
	}
)

var (
	iload0 = &ILoad{1}
	iload1 = &ILoad{2}
	iload2 = &ILoad{3}
	iload3 = &ILoad{4}
)

func (i *ILoad) Clone() instructions.Inst {
	if i.idx > 0 {
		return i
	}
	return &ILoad{}
}

func (i *ILoad) Fetch(coder *instructions.CodeReader) {
	if i.idx > 0 {
		return
	}
	i.idx = int(coder.Read1()) + 1
}

func (i *ILoad) Exec(f *stack.Frame) {
	iload(f, i.idx-1)
}

func iload(f *stack.Frame, idx int)  {
	f.PushOpstackVal(f.GetLocalVal(idx))
}

func init() {
	instructions.Register(iload_op, &ILoad{})
	instructions.Register(iload0_op, iload0)
	instructions.Register(iload1_op, iload1)
	instructions.Register(iload2_op, iload2)
	instructions.Register(iload3_op, iload3)
}
