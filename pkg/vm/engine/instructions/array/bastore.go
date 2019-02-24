package array

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	bastore_op = 0x54
)

type (
	BAstoreInst struct {
		
	}
)

func (i *BAstoreInst) Clone() instructions.Inst {
	return i
}

func (i *BAstoreInst) Fetch(coder *instructions.CodeReader) {

}

func (i *BAstoreInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction bastore")
	v := f.PopOpstackVal()
	idx := f.PopOpstackVal()
	arr := f.PopOpstackRef()
	arr.Bytes()[idx] = int8(v)
}

func init()  {
	instructions.Register(bastore_op, &BAstoreInst{})
}