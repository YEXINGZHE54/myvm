package objects

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	new_op = 0xbb
)

type (
	NewInst struct {
		idx uint16
	}
)

func (i *NewInst) Clone() instructions.Inst {
	return &NewInst{}
}

func (i *NewInst) Fetch(coder *instructions.CodeReader) {
	i.idx = coder.Read2()
}

func (i *NewInst) Exec(f *stack.Frame) {
	println("new op exec: ")
	cls := f.GetMethod().Cls
	ref := cls.Consts[i.idx].(*reflect.ClsRef)
	err := cls.Loader.ResolveClass(ref)
	if err != nil {
		panic(err)
	}
	f.PushOpstackRef(&reflect.Object{})
}

func init() {
	instructions.Register(new_op, &NewInst{})
}
