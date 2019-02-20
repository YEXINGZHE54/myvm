package branches

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	ifeq_op = 0x99
	ifne_op = 0x9a
	iflt_op = 0x9b
	ifge_op = 0x9c
	ifgt_op = 0x9d
	ifle_op = 0x9e
)

type (
	cond func(v int32) bool
	IfCondInst struct {
		f cond
		idx int16
	}
)

func (i *IfCondInst) Clone() instructions.Inst {
	return &IfCondInst{f: i.f} //clone condtion
}

func (i *IfCondInst) Fetch(coder *instructions.CodeReader) {
	i.idx = int16(coder.Read2())
}

func (i *IfCondInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction ifcond")
	if i.f(f.PopOpstackVal()) {
		gotoOffset(f, i.idx)
	}
}

func eq(v int32) bool {
	return v == 0
}

func ne(v int32) bool {
	return v != 0
}

func lt(v int32) bool {
	return v < 0
}

func gt(v int32) bool {
	return v > 0
}

func le(v int32) bool {
	return v <= 0
}

func ge(v int32) bool {
	return v >= 0
}

func init() {
	instructions.Register(ifeq_op, &IfCondInst{eq, 0})
	instructions.Register(ifne_op, &IfCondInst{ne, 0})
	instructions.Register(iflt_op, &IfCondInst{lt, 0})
	instructions.Register(ifge_op, &IfCondInst{ge, 0})
	instructions.Register(ifgt_op, &IfCondInst{gt, 0})
	instructions.Register(ifle_op, &IfCondInst{le, 0})
}