package convert

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	checkcast_op = 0xc0
)

type (
	CheckcastInst struct {
		idx uint16
	}
)

func (i *CheckcastInst) Clone() instructions.Inst {
	return &CheckcastInst{}
}

func (i *CheckcastInst) Fetch(coder *instructions.CodeReader)  {
	i.idx = coder.Read2()
}

func (i *CheckcastInst) Exec(f *stack.Frame)  {
	obj := f.GetOpstackSlot(0).Ref
	if obj == nil {
		return
	}
	tcls := f.GetMethod().Cls.Consts[i.idx].(*reflect.ClsRef)
	err := f.GetMethod().Cls.Loader.ResolveClass(tcls)
	if err != nil {
		panic(err)
	}
	scls := obj.Class
	if !reflect.CanCastTo(scls, tcls.Ref) {
		panic(" ClassCastException")
	}
}

func init()  {
	instructions.Register(checkcast_op, &CheckcastInst{})
}