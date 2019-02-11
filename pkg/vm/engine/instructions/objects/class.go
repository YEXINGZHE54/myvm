package objects

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	getstatic_op = 0xb2
	putstatic_op = 0xb3
)

type (
	GetStaticInst struct {
		idx uint16
	}
	PutStaticInst struct {
		idx uint16
	}
)

func (i *GetStaticInst) Clone() instructions.Inst {
	return &GetStaticInst{}
}

func (i *GetStaticInst) Fetch(coder *instructions.CodeReader) {
	i.idx = coder.Read2()
}

func (i *GetStaticInst) Exec(f *stack.Frame) {
	println("getstatic exec: ")
	f.PopOpstackRef() //TODO: somewhere error
	cls := f.GetMethod().Cls
	ref := cls.Consts[i.idx].(*reflect.FieldRef)
	err := cls.Loader.ResolveField(ref)
	if err != nil {
		panic(err)
	}
	switch ref.Ref.Desc[0] {
	case 'Z','B','C','S','I':
		f.PushOpstackVal(ref.Ref.Cls.StaticVars.GetVal(ref.Ref.SlotId))
	case 'F':
		f.PushOpstackFloat(ref.Ref.Cls.StaticVars.GetFloat(ref.Ref.SlotId))
	case 'J':
		f.PushOpstackLong(ref.Ref.Cls.StaticVars.GetLong(ref.Ref.SlotId))
	case 'D':
		f.PushOpstackDouble(ref.Ref.Cls.StaticVars.GetDouble(ref.Ref.SlotId))
	case 'L','[': //object or array
		f.PushOpstackRef(ref.Ref.Cls.StaticVars.GetRef(ref.Ref.SlotId))
	}
}

func (i *PutStaticInst) Clone() instructions.Inst {
	return &PutStaticInst{}
}

func (i *PutStaticInst) Fetch(coder *instructions.CodeReader) {
	i.idx = coder.Read2()
}

func (i *PutStaticInst) Exec(f *stack.Frame) {
	println("putstatic exec: ")
	cls := f.GetMethod().Cls
	ref := cls.Consts[i.idx].(*reflect.FieldRef)
	err := cls.Loader.ResolveField(ref)
	if err != nil {
		panic(err)
	}
	switch ref.Ref.Desc[0] {
	case 'Z','B','C','S','I':
		ref.Ref.Cls.StaticVars.SetVal(ref.Ref.SlotId, f.PopOpstackVal())
	case 'F':
		ref.Ref.Cls.StaticVars.SetFloat(ref.Ref.SlotId, f.PopOpstackFloat())
	case 'J':
		ref.Ref.Cls.StaticVars.SetLong(ref.Ref.SlotId, f.PopOpstackLong())
	case 'D':
		ref.Ref.Cls.StaticVars.SetDouble(ref.Ref.SlotId, f.PopOpstackDouble())
	case 'L','[': //object or array
		ref.Ref.Cls.StaticVars.SetRef(ref.Ref.SlotId, f.PopOpstackRef())
	}
}

func init() {
	instructions.Register(getstatic_op, &GetStaticInst{})
	instructions.Register(putstatic_op, &PutStaticInst{})
}