package array

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	caload_op = 0x34
)

type (
	CALoad struct{
	}
)

func (i *CALoad) Clone() instructions.Inst {
	return i
}

func (i *CALoad) Fetch(coder *instructions.CodeReader) {

}

func (i *CALoad) Exec(f *stack.Frame) {
	utils.Log("executing instruction calod")
	idx := int(f.PopOpstackVal())
	arr := f.PopOpstackRef()
	f.PushOpstackVal(int32(arr.Chars()[idx]))
}

func init()  {
	instructions.Register(caload_op, &CALoad{})
}