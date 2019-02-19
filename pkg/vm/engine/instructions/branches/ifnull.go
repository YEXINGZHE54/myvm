package branches

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	ifnull_op = 0xc6
)

type (
	IfNullInst struct {
		idx int16
	}
)

func (i *IfNullInst) Clone() instructions.Inst {
	return &IfNullInst{}
}

func (i *IfNullInst) Fetch(coder *instructions.CodeReader) {
	i.idx = int16(coder.Read2())
}

func (i *IfNullInst) Exec(f *stack.Frame) {
	o := f.PopOpstackRef()
	if o == nil {
		gotoOffset(f, i.idx)
	}
}

func init() {
	instructions.Register(ifnull_op, &IfNullInst{})
}