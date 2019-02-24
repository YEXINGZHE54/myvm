package branches

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	if_acmpeq_op = 0xa5
	if_acmpne_op = 0xa6
)

type (
	acmpcond func(v1, v2 *reflect.Object) bool
	IfAcmpInst struct {
		f acmpcond
		idx int16
	}
)

func (i *IfAcmpInst) Clone() instructions.Inst {
	return &IfAcmpInst{i.f, 0} //clone cond
}

func (i *IfAcmpInst) Fetch(coder *instructions.CodeReader) {
	i.idx = int16(coder.Read2())
}

func (i *IfAcmpInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction ifacmp")
	v2 := f.PopOpstackRef()
	v1 := f.PopOpstackRef()
	if i.f(v1, v2) {
		gotoOffset(f, i.idx)
	}
}

func acmp_eq(v1, v2 *reflect.Object) bool {
	return v1 == v2
}

func acmp_ne(v1, v2 *reflect.Object) bool {
	return v1 != v2
}

func init()  {
	instructions.Register(if_acmpeq_op, &IfAcmpInst{acmp_eq, 0})
	instructions.Register(if_acmpne_op, &IfAcmpInst{acmp_ne, 0})
}