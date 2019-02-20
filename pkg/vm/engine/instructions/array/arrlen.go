package array

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	arrlen_op = 0xbe
)

type (
	ArrayLenInst struct{
	}
)

func (i *ArrayLenInst) Clone() instructions.Inst {
	return i
}

func (i *ArrayLenInst) Fetch(coder *instructions.CodeReader) {

}

func (i *ArrayLenInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction arrlen")
	f.PushOpstackVal(int32(f.PopOpstackRef().ArrLength()))
}

func init() {
	instructions.Register(arrlen_op, &ArrayLenInst{})
}
