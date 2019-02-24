package array

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	baload_op = 0x33
)

type (
	BALoad struct{
	}
)

func (i *BALoad) Clone() instructions.Inst {
	return i
}

func (i *BALoad) Fetch(coder *instructions.CodeReader) {

}

func (i *BALoad) Exec(f *stack.Frame) {
	utils.Log("executing instruction balod")
	idx := int(f.PopOpstackVal())
	arr := f.PopOpstackRef()
	f.PushOpstackVal(int32(arr.Bytes()[idx]))
}

func init()  {
	instructions.Register(baload_op, &BALoad{})
}