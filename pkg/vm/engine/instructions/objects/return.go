package objects

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	return_op = 0xb1
	ireturn_op = 0xac
	lreturn_op = 0xad
	freturn_op = 0xae
	dreturn_op = 0xaf
	areturn_op = 0xb0
)

type (
	ReturnInst struct {}
	IReturnInst struct{}
	LReturnInst struct{}
	FReturnInst struct{}
	DReturnInst struct{}
	AReturnInst struct{}
)

func (i *ReturnInst) Clone() instructions.Inst {
	return i
}

func (i *ReturnInst) Fetch(coder *instructions.CodeReader) {

}

func (i *ReturnInst) Exec(f *stack.Frame) {
	println("return exec")
	f.Stack.Pop()
}

func (i *IReturnInst) Clone() instructions.Inst {
	return i
}

func (i *IReturnInst) Fetch(coder *instructions.CodeReader) {

}

func (i *IReturnInst) Exec(f *stack.Frame) {
	println("ireturn exec")
	top := f.Stack.Pop()
	ret := f.Stack.Current()
	ret.PushOpstackVal(top.PopOpstackVal())
}

func (i *LReturnInst) Clone() instructions.Inst {
	return i
}

func (i *LReturnInst) Fetch(coder *instructions.CodeReader) {

}

func (i *LReturnInst) Exec(f *stack.Frame) {
	println("lreturn exec")
	top := f.Stack.Pop()
	ret := f.Stack.Current()
	ret.PushOpstackLong(top.PopOpstackLong())
}

func (i *FReturnInst) Clone() instructions.Inst {
	return i
}

func (i *FReturnInst) Fetch(coder *instructions.CodeReader) {

}

func (i *FReturnInst) Exec(f *stack.Frame) {
	println("freturn exec")
	top := f.Stack.Pop()
	ret := f.Stack.Current()
	ret.PushOpstackFloat(top.PopOpstackFloat())
}

func (i *DReturnInst) Clone() instructions.Inst {
	return i
}

func (i *DReturnInst) Fetch(coder *instructions.CodeReader) {

}

func (i *DReturnInst) Exec(f *stack.Frame) {
	println("dreturn exec")
	top := f.Stack.Pop()
	ret := f.Stack.Current()
	ret.PushOpstackDouble(top.PopOpstackDouble())
}

func (i *AReturnInst) Clone() instructions.Inst {
	return i
}

func (i *AReturnInst) Fetch(coder *instructions.CodeReader) {

}

func (i *AReturnInst) Exec(f *stack.Frame) {
	println("areturn exec")
	top := f.Stack.Pop()
	ret := f.Stack.Current()
	ret.PushOpstackRef(top.PopOpstackRef())
}

func init() {
	instructions.Register(return_op, &ReturnInst{})
	instructions.Register(ireturn_op, &IReturnInst{})
	instructions.Register(lreturn_op, &LReturnInst{})
	instructions.Register(freturn_op, &FReturnInst{})
	instructions.Register(dreturn_op, &DReturnInst{})
	instructions.Register(areturn_op, &AReturnInst{})
}