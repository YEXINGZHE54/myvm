package array

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	aastore_op = 0x53
)

type (
	AAstoreInst struct {
		
	}
)

func (i *AAstoreInst) Clone() instructions.Inst {
	return i
}

func (i *AAstoreInst) Fetch(coder *instructions.CodeReader) {

}

func (i *AAstoreInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction aastore")
	v := f.PopOpstackRef()
	idx := f.PopOpstackVal()
	arr := f.PopOpstackRef()
	arr.Refs()[idx] = v
}

func init()  {
	instructions.Register(aastore_op, &AAstoreInst{})
}