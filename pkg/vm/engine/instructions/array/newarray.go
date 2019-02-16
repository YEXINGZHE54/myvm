package array

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	newarray_op = 0xbc
	anewarray_op = 0xbd
	ATYPE_BOOLEAN = 4
	ATYPE_CHAR = 5
	ATYPE_FLOAT = 6
	ATYPE_DOUBLE = 7
	ATYPE_BYTE = 8
	ATYPE_SHORT = 9
	ATYPE_INT = 10
	ATYPE_LONG = 11
)

type (
	NewArrayInst struct{
		idx uint8
	}
	ANewArrayInst struct{
		idx uint16
	}
)

func (i *NewArrayInst) Clone() instructions.Inst {
	return &NewArrayInst{}
}

func (i *NewArrayInst) Fetch(coder *instructions.CodeReader) {
	i.idx = coder.Read1()
}

func (i *NewArrayInst) Exec(f *stack.Frame) {
	var cname string
	switch i.idx {
	case ATYPE_BOOLEAN:
		cname = "[B"
	case ATYPE_CHAR:
		cname = "[C"
	case ATYPE_FLOAT:
		cname = "[F"
	case ATYPE_DOUBLE:
		cname = "[D"
	case ATYPE_BYTE:
		cname = "[B"
	case ATYPE_SHORT:
		cname = "[S"
	case ATYPE_INT:
		cname = "[I"
	case ATYPE_LONG:
		cname = "[J"
	default:
		panic("unexpected array type")
	}
	count := f.PopOpstackVal()
	arrcls, err := f.GetMethod().Cls.Loader.LoadClass(cname)
	if err != nil {
		panic(err)
	}
	o, err := arrcls.NewArray(int(count))
	if err != nil {
		panic(err)
	}
	f.PushOpstackRef(o)
}

func (i *ANewArrayInst) Clone() instructions.Inst {
	return &ANewArrayInst{}
}

func (i *ANewArrayInst) Fetch(coder *instructions.CodeReader) {
	i.idx = coder.Read2()
}

func (i *ANewArrayInst) Exec(f *stack.Frame) {
	ref := f.GetMethod().Cls.Consts[i.idx].(*reflect.ClsRef)
	err := f.GetMethod().Cls.Loader.ResolveClass(ref)
	if err != nil {
		panic(err)
	}
	arrcls, err := ref.Ref.ArrayClass()
	if err != nil {
		panic(err)
	}
	count := f.PopOpstackVal()
	o, err := arrcls.NewArray(int(count))
	if err != nil {
		panic(err)
	}
	f.PushOpstackRef(o)
}

func init() {
	instructions.Register(newarray_op, &NewArrayInst{})
	instructions.Register(anewarray_op, &ANewArrayInst{})
}
