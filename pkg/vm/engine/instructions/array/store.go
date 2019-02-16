package array

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	iastore_op = 0x4f
)

type (
	IArrStore struct{
	}
)

func (i *IArrStore) Clone() instructions.Inst {
	return i
}

func (i *IArrStore) Fetch(coder *instructions.CodeReader) {

}

func (i *IArrStore) Exec(f *stack.Frame) {
	val := f.PopOpstackVal()
	idx := int(f.PopOpstackVal())
	arr := f.PopOpstackRef() // must be of type [I
	arr.Ints()[idx] = val
}

func init() {
	instructions.Register(iastore_op, &IArrStore{})
}
