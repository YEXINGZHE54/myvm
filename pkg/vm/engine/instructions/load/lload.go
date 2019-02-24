package load

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	lload_op = 0x16
	lload0_op = 0x1e
	lload1_op = 0x1f
	lload2_op = 0x20
	lload3_op = 0x21
)

type (
	LLoad struct {
		idx int
	}
)

var (
	lload0 = &LLoad{1}
	lload1 = &LLoad{2}
	lload2 = &LLoad{3}
	lload3 = &LLoad{4}
)

func (i *LLoad) Clone() instructions.Inst {
	if i.idx > 0 {
		return i
	}
	return &LLoad{}
}

func (i *LLoad) Fetch(coder *instructions.CodeReader) {
	if i.idx > 0 {
		return
	}
	i.idx = int(coder.Read1()) + 1
}

func (i *LLoad) Exec(f *stack.Frame) {
	utils.Log("executing instruction lload")
	lload(f, i.idx-1)
}

func lload(f *stack.Frame, idx int)  {
	f.PushOpstackLong(f.GetLocalLong(idx))
}

func init() {
	instructions.Register(lload_op, &LLoad{})
	instructions.Register(lload0_op, lload0)
	instructions.Register(lload1_op, lload1)
	instructions.Register(lload2_op, lload2)
	instructions.Register(lload3_op, lload3)
}
