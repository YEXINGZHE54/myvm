package array

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	iastore_op = 0x4f
	castore_op = 0x55
)

type (
	IArrStore struct{
	}
	CArrStore struct {
	}
)

func (i *IArrStore) Clone() instructions.Inst {
	return i
}

func (i *IArrStore) Fetch(coder *instructions.CodeReader) {

}

func (i *IArrStore) Exec(f *stack.Frame) {
	utils.Log("executing instruction iarrstore")
	val := f.PopOpstackVal()
	idx := int(f.PopOpstackVal())
	arr := f.PopOpstackRef() // must be of type [I
	arr.Ints()[idx] = val
}

func (i *CArrStore) Clone() instructions.Inst {
	return i
}

func (i *CArrStore) Fetch(coder *instructions.CodeReader) {

}

func (i *CArrStore) Exec(f *stack.Frame) {
	utils.Log("executing instruction carrstore")
	val := f.PopOpstackVal()
	idx := int(f.PopOpstackVal())
	arr := f.PopOpstackRef() // must be of type [I
	arr.Chars()[idx] = uint16(val)
}

func init() {
	instructions.Register(iastore_op, &IArrStore{})
	instructions.Register(castore_op, &CArrStore{})
}
