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
	invokeinterface_op = 0xb9
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
	InvokeInterfaceInst struct {
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

/*
1. If the resolved method is protected, and it is a member of a superclass of the current class,
and the method is not declared in the same run-time package (§5.3) as the current class,
then the class of objectref must be either the current class or a subclass of the current class.
2.1. The resolved method is not an instance initialization method,
and If the symbolic reference names a class (not an interface), then that class is a superclass of the current class,
and The ACC_SUPER flag is set for the class file
2.2 then C = direct super of current Class, otherwize C = method resolved Class/Interface,
then invokeMethod = lookupSpecialMethod(C, resolvedMethdo)
3. If the method is synchronized,
the monitor associated with objectref is entered or reentered
as if by execution of a monitorenter instruction (§monitorenter) in the current thread.
 */
func (i *InvokeSpecialInst) Exec(f *stack.Frame) {
	println("invoke special exec")
	cls := f.GetMethod().Cls
	ref := cls.Consts[i.idx].(*reflect.MethodRef)
	err := cls.Loader.ResolveMethod(ref)
	if err != nil {
		panic(err)
	}
	// 1
	m := ref.Ref
	if m.IsProtected() &&
		m.Cls.IsSuperOf(cls) &&
		m.Cls.GetPackageName() != cls.GetPackageName() &&
		m.Cls != cls &&
		!cls.IsSuperOf(m.Cls) {
		panic("protected method class must be a subclass of current class")
	}
	// 2.1
	var c *reflect.Class
	if m.Name != "<init>" &&
		m.Cls.IsClass() &&
		m.Cls.IsSuperOf(cls) &&
		cls.IsSuperSet() {
		c = cls.Super
	} else {
		c = m.Cls
	}
	// 2.2
	invoked, err := c.LookupSpecialMethod(ref.Ref)
	if err != nil {
		panic(err)
	}
	// 3 sync?
	invokeMethod(f, invoked)
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
	// check class init
	inited, err := init_class(f, ref.Ref.Cls)
	if err != nil {
		panic(err)
	}
	if !inited {
		revertPC(f)
		return
	}
	invokeMethod(f, ref.Ref)
}


func (i *InvokeInterfaceInst) Clone() instructions.Inst {
	return &InvokeInterfaceInst{}
}

func (i *InvokeInterfaceInst) Fetch(coder *instructions.CodeReader) {
	i.idx = coder.Read2()
	coder.Read2()
}

func (i *InvokeInterfaceInst) Exec(f *stack.Frame) {
	println("invoke interface exec")
	println(i.idx)
	//f.PushOpstackRef(nil)
}

func init() {
	instructions.Register(invokevirtual_op, &InvokeVirtualInst{})
	instructions.Register(invokespecial_op, &InvokeVirtualInst{})
	instructions.Register(invokestatic_op, &InvokeStaticInst{})
	instructions.Register(invokeinterface_op, &InvokeInterfaceInst{})
}