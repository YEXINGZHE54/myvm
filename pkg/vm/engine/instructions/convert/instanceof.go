package convert

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	instanceof_op = 0xc1
)

type (
	InstanceofInst struct {
		idx uint16
	}
)

func (i *InstanceofInst) Clone() instructions.Inst {
	return &InstanceofInst{}
}

func (i *InstanceofInst) Fetch(coder *instructions.CodeReader) {
	i.idx = coder.Read2()
}

func (i *InstanceofInst) Exec(f *stack.Frame) {
	o := f.PopOpstackRef()
	if o == nil {
		f.PushOpstackVal(0)
		return
	}
	clsref := f.GetMethod().Cls.Consts[i.idx].(*reflect.ClsRef)
	err := f.GetMethod().Cls.Loader.ResolveClass(clsref)
	if err != nil {
		panic(err)
	}
	if reflect.CanCastTo(o.Class, clsref.Ref) {
		f.PushOpstackVal(1)
	} else {
		f.PushOpstackVal(0)
	}
}

func init()  {
	instructions.Register(instanceof_op, &InstanceofInst{})
}