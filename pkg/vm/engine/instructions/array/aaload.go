package array

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	aaload_op = 0x32
)

type (
	AALoad struct{
	}
)

func (i *AALoad) Clone() instructions.Inst {
	return i
}

func (i *AALoad) Fetch(coder *instructions.CodeReader) {

}

func (i *AALoad) Exec(f *stack.Frame) {
	idx := int(f.PopOpstackVal())
	arr := f.PopOpstackRef()
	f.PushOpstackRef(arr.Refs()[idx])
}

func init()  {
	instructions.Register(aaload_op, &AALoad{})
}