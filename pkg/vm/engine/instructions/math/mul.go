package math

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	fmul_op = 0x6a
	imul_op = 0x68
)

type (
	FMulInst struct {
	}
	IMulInst struct {
	}
)

func (i *FMulInst) Clone() instructions.Inst {
	return i
}

func (i *FMulInst) Fetch(coder *instructions.CodeReader)  {

}

func (i *FMulInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction fmul")
	v2 := f.PopOpstackFloat()
	v1 := f.PopOpstackFloat()
	f.PushOpstackFloat(v1 * v2)
}

func (i *IMulInst) Clone() instructions.Inst {
	return i
}

func (i *IMulInst) Fetch(coder *instructions.CodeReader)  {

}

func (i *IMulInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction imul")
	v2 := f.PopOpstackVal()
	v1 := f.PopOpstackVal()
	f.PushOpstackVal(v1 * v2)
}

func init()  {
	instructions.Register(fmul_op, &FMulInst{})
	instructions.Register(imul_op, &IMulInst{})
}