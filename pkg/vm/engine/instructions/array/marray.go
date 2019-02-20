package array

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	multianewarray_op = 0xc5
)

type (
	MultiArrayInst struct{
		idx uint16
		dim uint8
	}
)

func (i *MultiArrayInst) Clone() instructions.Inst {
	return &MultiArrayInst{}
}

func (i *MultiArrayInst) Fetch(coder *instructions.CodeReader) {
	i.idx = coder.Read2()
	i.dim = coder.Read1()
}

func (i *MultiArrayInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction marray")
	counts := make([]int, i.dim)
	for idx := int(i.dim)-1; idx >= 0; idx = idx - 1 {
		counts[idx] = int(f.PopOpstackVal())
	}
	ref := f.GetMethod().Cls.Consts[i.idx].(*reflect.ClsRef)
	err := f.GetMethod().Cls.Loader.ResolveClass(ref)
	if err != nil {
		panic(err)
	}
	arrcls := ref.Ref
	o, err := arrcls.NewMArray(counts)
	if err != nil {
		panic(err)
	}
	f.PushOpstackRef(o)
}

func init() {
	instructions.Register(multianewarray_op, &MultiArrayInst{})
}