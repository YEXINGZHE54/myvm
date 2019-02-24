package load

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	dload_op = 0x18
	dload0_op = 0x26
	dload1_op = 0x27
	dload2_op = 0x28
	dload3_op = 0x29
)

type (
	DLoad struct {
		idx int
	}
)

var (
	dload0 = &DLoad{1}
	dload1 = &DLoad{2}
	dload2 = &DLoad{3}
	dload3 = &DLoad{4}
)

func (i *DLoad) Clone() instructions.Inst {
	if i.idx > 0 {
		return i
	}
	return &DLoad{}
}

func (i *DLoad) Fetch(coder *instructions.CodeReader) {
	if i.idx > 0 {
		return
	}
	i.idx = int(coder.Read1()) + 1
}

func (i *DLoad) Exec(f *stack.Frame) {
	utils.Log("executing instruction dload")
	dload(f, i.idx-1)
}

func dload(f *stack.Frame, idx int)  {
	f.PushOpstackDouble(f.GetLocalDouble(idx))
}

func init() {
	instructions.Register(dload_op, &DLoad{})
	instructions.Register(dload0_op, dload0)
	instructions.Register(dload1_op, dload1)
	instructions.Register(dload2_op, dload2)
	instructions.Register(dload3_op, dload3)
}
