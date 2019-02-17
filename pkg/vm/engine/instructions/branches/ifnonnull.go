package branches

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	ifnonnull_op = 0xc7
)

type (
	IfNonNullInst struct {
		idx int16
	}
)

func (i *IfNonNullInst) Clone() instructions.Inst {
	return &IfNonNullInst{}
}

func (i *IfNonNullInst) Fetch(coder *instructions.CodeReader) {
	i.idx = int16(coder.Read2())
}

func (i *IfNonNullInst) Exec(f *stack.Frame) {
	o := f.PopOpstackRef()
	if o != nil {
		gotoOffset(f, i.idx)
	}
}

func init() {
	instructions.Register(ifnonnull_op, &IfNonNullInst{})
}