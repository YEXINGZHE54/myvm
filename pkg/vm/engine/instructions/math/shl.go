package math

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	lshl_op = 0x79
	ishl_op = 0x78
)

type (
	lshlInst struct {
	}
	ishlInst struct {
	}
)

func (i *lshlInst) Clone() instructions.Inst {
	return i
}

func (i *lshlInst) Fetch(coder *instructions.CodeReader) {

}

func (i *lshlInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction lshl")
	v2 := f.PopOpstackVal()
	v1 := f.PopOpstackLong()
	shift := uint(v2 & 0x3F)
	f.PushOpstackLong(v1 << shift)
}

func (i *ishlInst) Clone() instructions.Inst {
	return i
}

func (i *ishlInst) Fetch(coder *instructions.CodeReader) {

}

func (i *ishlInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction ishl")
	v2 := f.PopOpstackVal()
	v1 := f.PopOpstackVal()
	shift := uint(v2 & 0x3F)
	f.PushOpstackVal(v1 << shift)
}

func init()  {
	instructions.Register(lshl_op, &lshlInst{})
	instructions.Register(ishl_op, &ishlInst{})
}