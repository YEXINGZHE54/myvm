package array

import (
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
	idx := int(f.PopOpstackVal())
	arr := f.PopOpstackRef() // must be of type [I
	f.PushOpstackVal(int32(arr.Ints()[idx]))
}

func init() {
	instructions.Register(iaload_op, &IArrLoad{})
}
