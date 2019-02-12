package objects

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	invokevirtual_op = 0xb6
	invokespecial_op = 0xb7
	invokestatic_op = 0xb8
)

type (
	InvokeVirtualInst struct {
		idx uint16
	}
	InvokeSpecialInst struct {
		idx uint16
	}
	InvokeStaticInst struct {
		idx uint16
	}
)

func (i *InvokeVirtualInst) Clone() instructions.Inst {
	return &InvokeVirtualInst{}
}

func (i *InvokeVirtualInst) Fetch(coder *instructions.CodeReader) {
	i.idx = coder.Read2()
}

func (i *InvokeVirtualInst) Exec(f *stack.Frame) {
	println("invoke virtual exec")
	println(i.idx)
	//f.PushOpstackRef(nil)
}

func (i *InvokeSpecialInst) Clone() instructions.Inst {
	return &InvokeSpecialInst{}
}

func (i *InvokeSpecialInst) Fetch(coder *instructions.CodeReader) {
	i.idx = coder.Read2()
}

func (i *InvokeSpecialInst) Exec(f *stack.Frame) {
	println("invoke special exec")
	cls := f.GetMethod().Cls
	ref := cls.Consts[i.idx].(*reflect.MethodRef)
	err := cls.Loader.ResolveMethod(ref)
	if err != nil {
		panic(err)
	}
	if ref.Ref.IsStatic() {
		panic("not expecting static method in invokespecial op")
	}
}

func (i *InvokeStaticInst) Clone() instructions.Inst {
	return &InvokeStaticInst{}
}

func (i *InvokeStaticInst) Fetch(coder *instructions.CodeReader) {
	i.idx = coder.Read2()
}

func (i *InvokeStaticInst) Exec(f *stack.Frame) {
	println("invoke static exec")
	cls := f.GetMethod().Cls
	ref := cls.Consts[i.idx].(*reflect.MethodRef)
	err := cls.Loader.ResolveMethod(ref)
	if err != nil {
		if ref.Name == "registerNatives" { //TODO: skip native
			return
		}
		panic(err)
	}
	if !ref.Ref.IsStatic() {
		panic("expecting static method in invokestatic op")
	}
	invokeMethod(f, ref.Ref)
}

func init() {
	instructions.Register(invokevirtual_op, &InvokeVirtualInst{})
	instructions.Register(invokespecial_op, &InvokeVirtualInst{})
	instructions.Register(invokestatic_op, &InvokeStaticInst{})
}