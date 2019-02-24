package array

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	iaload_op = 0x2e
)

type (
	IArrLoad struct{
	}
)

func (i *IArrLoad) Clone() instructions.Inst {
	return i
}

func (i *IArrLoad) Fetch(coder *instructions.CodeReader) {

}

func (i *IArrLoad) Exec(f *stack.Frame) {
	utils.Log("executing instruction iaload")
	idx := int(f.PopOpstackVal())
	arr := f.PopOpstackRef() // must be of type [I
	f.PushOpstackVal(arr.Ints()[idx])
}

func init() {
	instructions.Register(iaload_op, &IArrLoad{})
}
