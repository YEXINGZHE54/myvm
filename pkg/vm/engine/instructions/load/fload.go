package load

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	fload_op = 0x17
	fload0_op = 0x22
	fload1_op = 0x23
	fload2_op = 0x24
	fload3_op = 0x25
)

type (
	FLoad struct {
		idx int
	}
)

var (
	fload0 = &FLoad{1}
	fload1 = &FLoad{2}
	fload2 = &FLoad{3}
	fload3 = &FLoad{4}
)

func (i *FLoad) Clone() instructions.Inst {
	if i.idx > 0 {
		return i
	}
	return &FLoad{}
}

func (i *FLoad) Fetch(coder *instructions.CodeReader) {
	if i.idx > 0 {
		return
	}
	i.idx = int(coder.Read1()) + 1
}

func (i *FLoad) Exec(f *stack.Frame) {
	utils.Log("executing instruction fload")
	fload(f, i.idx-1)
}

func fload(f *stack.Frame, idx int)  {
	f.PushOpstackVal(f.GetLocalVal(idx))
}

func init() {
	instructions.Register(fload_op, &FLoad{})
	instructions.Register(fload0_op, fload0)
	instructions.Register(fload1_op, fload1)
	instructions.Register(fload2_op, fload2)
	instructions.Register(fload3_op, fload3)
}
