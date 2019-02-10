package constants

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/loader/classfile"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
)

const (
	ldc_op = 0x12
	ldcw_op = 0x13
	ldc2w_op = 0x14
)

type (
	LdcInst struct {
		idx uint8
	}
	LdcwInst struct {
		idx uint16
	}
	Ldc2wInst struct {
		idx uint16
	}
)

func (i *LdcInst) Clone() instructions.Inst {
	return &LdcInst{}
}

func (i *LdcInst) Fetch(coder *classfile.CodeReader) {
	i.idx = coder.Read1()
}

func (i *LdcInst) Exec(f *stack.Frame) {
	println("ldc exec")
	println(i.idx)
	//f.PushOpstackRef(nil)
}

func (i *LdcwInst) Clone() instructions.Inst {
	return &LdcwInst{}
}

func (i *LdcwInst) Fetch(coder *classfile.CodeReader) {
	i.idx = coder.Read2()
}

func (i *LdcwInst) Exec(f *stack.Frame) {
	println("ldcw exec")
	//f.PushOpstackRef(nil)
}

func (i *Ldc2wInst) Clone() instructions.Inst {
	return &Ldc2wInst{}
}

func (i *Ldc2wInst) Fetch(coder *classfile.CodeReader) {
	i.idx = coder.Read2()
}

func (i *Ldc2wInst) Exec(f *stack.Frame) {
	println("ldc2w exec")
	//f.PushOpstackRef(nil)
}

func init() {
	instructions.Register(ldc_op, &LdcInst{})
	instructions.Register(ldcw_op, &LdcwInst{})
	instructions.Register(ldc2w_op, &Ldc2wInst{})
}