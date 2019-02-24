package math

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	iushr_op = 0x7c
	ishr_op = 0x7a
)

type (
	iushrInst struct {
	}
	ishrInst struct {
	}
)

func (i *iushrInst) Clone() instructions.Inst {
	return i
}

func (i *iushrInst) Fetch(coder *instructions.CodeReader) {

}

func (i *iushrInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction iushr")
	v2 := f.PopOpstackVal()
	v1 := f.PopOpstackVal()
	s := uint(v2 & 0x1F)
	if s >= 0 {
		f.PushOpstackVal(v1 >> s)
	} else {
		f.PushOpstackVal(v1 >> s + 2 << (31-s))
	}
}

func (i *ishrInst) Clone() instructions.Inst {
	return i
}

func (i *ishrInst) Fetch(coder *instructions.CodeReader) {

}

func (i *ishrInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction ishr")
	v2 := f.PopOpstackVal()
	v1 := f.PopOpstackVal()
	s := uint(v2 & 0x1F)
	f.PushOpstackVal(v1 >> s)
}

func init()  {
	instructions.Register(iushr_op, &iushrInst{})
	instructions.Register(ishr_op, &ishrInst{})
}