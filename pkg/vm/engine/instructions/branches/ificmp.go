package branches

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	if_icmpeq_op = 0x9f
	if_icmpne_op = 0xa0
	if_icmplt_op = 0xa1
	if_icmpge_op = 0xa2
	if_icmpgt_op = 0xa3
	if_icmple_op = 0xa4
)

type (
	cmpcond func(v1, v2 int32) bool
	IfIcmpInst struct {
		f cmpcond
		idx int16
	}
)

func (i *IfIcmpInst) Clone() instructions.Inst {
	return &IfIcmpInst{i.f, 0} //clone cond
}

func (i *IfIcmpInst) Fetch(coder *instructions.CodeReader) {
	i.idx = int16(coder.Read2())
}

func (i *IfIcmpInst) Exec(f *stack.Frame) {
	v2 := f.PopOpstackVal()
	v1 := f.PopOpstackVal()
	if i.f(v1, v2) {
		gotoOffset(f, i.idx)
	}
}

func cmp_eq(v1, v2 int32) bool {
	return v1 == v2
}

func cmp_ne(v1, v2 int32) bool {
	return v1 != v2
}

func cmp_lt(v1, v2 int32) bool {
	return v1 < v2
}

func cmp_ge(v1, v2 int32) bool {
	return v1 >= v2
}

func cmp_gt(v1, v2 int32) bool {
	return v1 > v2
}

func cmp_le(v1, v2 int32) bool {
	return v1 <= v2
}

func init()  {
	instructions.Register(if_icmpeq_op, &IfIcmpInst{cmp_eq, 0})
	instructions.Register(if_icmpne_op, &IfIcmpInst{cmp_ne, 0})
	instructions.Register(if_icmplt_op, &IfIcmpInst{cmp_lt, 0})
	instructions.Register(if_icmpge_op, &IfIcmpInst{cmp_ge, 0})
	instructions.Register(if_icmpgt_op, &IfIcmpInst{cmp_gt, 0})
	instructions.Register(if_icmple_op, &IfIcmpInst{cmp_le, 0})
}